package automation

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/dop251/goja"
	cron "github.com/robfig/cron/v3"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	longrunningv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1/longrunningv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/wellknown"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules/connect"
	"github.com/tierklinik-dobersberg/longrunning-service/pkg/op"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Broker interface {
	Publish(*eventsv1.Event) error
	Subscribe(string, chan *eventsv1.Event)
}

type CoreModule struct {
	engine *Engine
	broker Broker

	scheduler *cron.Cron
}

func NewCoreModule(engine *Engine, broker Broker) *CoreModule {
	log := cron.PrintfLogger(slog.NewLogLogger(slog.NewTextHandler(os.Stderr, nil), slog.LevelInfo))

	scheduler := cron.New(
		cron.WithLocation(time.Local),
		cron.WithChain(
			cron.Recover(log),
			cron.DelayIfStillRunning(log),
		),
	)

	cm := &CoreModule{
		engine:    engine,
		scheduler: scheduler,
		broker:    broker,
	}

	cm.scheduler.Start()

	return cm
}

func (c *CoreModule) Enable(r *goja.Runtime) {
	r.Set("schedule", c.schedule)
	r.Set("clearSchedule", c.clearSchedule)
	r.Set("on", c.onEvent)
	r.Set("publish", c.publish)
}

func (c *CoreModule) schedule(schedule string, callable goja.Callable) (int, error) {
	c.engine.log.Log(context.Background(), slog.LevelInfo, "automation: new schedule registered", "schedule", schedule)

	res, err := c.scheduler.AddFunc(schedule, func() {
		c.engine.log.Log(context.Background(), slog.LevelInfo, "triggering automation schedule", "schedule", schedule)

		c.wrapOperation(callable, "schedule:"+fmt.Sprintf("%q", schedule), nil)
	})
	if err != nil {
		return -1, err
	}

	return int(res), nil
}

func (c *CoreModule) wrapOperation(callable goja.Callable, kind string, this any, args ...any) {
	var cli longrunningv1connect.LongRunningServiceClient
	if c.engine.automationConfig.WrapInOperation {
		var err error

		cli, err = wellknown.LongRunningService.Create(context.Background(), c.engine.discoverer)
		if err != nil {
			c.engine.log.Log(context.Background(), slog.LevelError, "failed to get longrunning service instance", "error", err)
		}
	}

	if cli == nil {
		c.engine.loop.RunOnLoop(func(r *goja.Runtime) {
			this := r.ToValue(this)
			a := make([]goja.Value, len(args))
			for idx, arg := range args {
				a[idx] = r.ToValue(arg)
			}

			callable(this, a...)
		})
	} else {
		_, err := op.Wrap(context.Background(), cli, func(ctx context.Context) (any, error) {
			id, _ := op.IDFromContext(ctx)
			token, _ := op.AuthTokenFromContext(ctx)

			stream := cli.StreamOperationLog(ctx)
			defer stream.CloseAndReceive()

			var prevLog *logger
			opLogger := &modules.FunctionLogger{
				Handler: func(ctx context.Context, l slog.Level, s string, a ...any) {
					buf := new(bytes.Buffer)
					h := slog.NewTextHandler(buf, nil)

					r := slog.NewRecord(time.Now(), l, s, 0)
					r.Add(args...)

					h.Handle(ctx, r)

					var severity longrunningv1.Severity
					switch l {
					case slog.LevelDebug:
						severity = longrunningv1.Severity_SEVERITY_DEBUG
					case slog.LevelInfo:
						severity = longrunningv1.Severity_SEVERITY_INFO
					case slog.LevelWarn:
						severity = longrunningv1.Severity_SEVERITY_WARNING
					case slog.LevelError:
						severity = longrunningv1.Severity_SEVERITY_ERROR
					}

					err := stream.Send(&longrunningv1.StreamOperationLogRequest{
						UniqueId:  id,
						AuthToken: token,
						Logs: []*longrunningv1.OperationLog{
							{
								Time:     timestamppb.Now(),
								Message:  buf.String(),
								Severity: severity,
							},
						},
					})

					if err != nil && prevLog != nil {
						prevLog.Log(ctx, slog.LevelError, "failed to send log message to long-running service", "error", err)
					}
				},
			}

			prevLog = c.engine.log.AddLogger(opLogger)
			defer c.engine.log.RemoveLogger(opLogger)

			var (
				result    any
				resultErr error
			)

			c.engine.log.Log(context.Background(), slog.LevelInfo, "scheduling operation on event loop")

			var wg sync.WaitGroup
			wg.Add(1)

			c.engine.loop.RunOnLoop(func(r *goja.Runtime) {
				defer wg.Done()

				this := r.ToValue(this)
				a := make([]goja.Value, len(args))
				for idx, arg := range args {
					a[idx] = r.ToValue(arg)
				}

				gv, err := callable(this, a...)

				if err == nil {
					result = gv.Export()
				} else {
					resultErr = err
				}
			})

			wg.Wait()

			return result, resultErr
		}, func(req *connect_go.Request[longrunningv1.RegisterOperationRequest]) {
			req.Msg.Kind = kind
			req.Msg.Owner = "automation"
			req.Msg.Description = c.engine.name

			cfg := c.engine.AutomationConfig()
			if len(cfg.ConnectHeaders) > 0 {
				for key, value := range cfg.ConnectHeaders {
					req.Header().Add(key, value)
				}
			}
		})

		if err != nil {
			c.engine.log.Log(context.Background(), slog.LevelError, "failed to execute goja callable", "error", err)
		}
	}
}

func (c *CoreModule) clearSchedule(id int) {
	c.scheduler.Remove(cron.EntryID(id))
}

func (c *CoreModule) publish(typeUrl string, obj *goja.Object) error {
	d, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(typeUrl))
	if err != nil {
		return err
	}

	md, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		return fmt.Errorf("invalid type url")
	}

	msg, err := connect.ObjectToProto(obj, md, c.engine.resolver)
	if err != nil {
		return err
	}

	evt, err := anypb.New(msg)
	if err != nil {
		return err
	}

	return c.broker.Publish(&eventsv1.Event{
		Event: evt,
	})
}

func (c *CoreModule) onEvent(event string, callable goja.Callable) {
	msgs := make(chan *eventsv1.Event, 100)

	c.broker.Subscribe(event, msgs)

	c.engine.log.Log(context.Background(), slog.LevelInfo, "automation: script successfully subscribed to event topic", "event", event)

	go func() {
		ctx := context.Background()
		defer c.engine.log.Log(ctx, slog.LevelInfo, "automation: subscription loop closed")
		for m := range msgs {
			c.engine.log.Log(ctx, slog.LevelInfo, "automation: received event, converting from proto-message", "typeUrl", m.Event.TypeUrl)

			o, err := connect.ConvertProtoMessage(m, c.engine.resolver)
			if err != nil {
				c.engine.log.Log(ctx, slog.LevelError, "failed to convert protobuf message", "error", err)
				continue
			}

			c.engine.log.Log(ctx, slog.LevelInfo, "running automation for event", "typeUrl", m.Event.TypeUrl)

			c.wrapOperation(callable, "event:"+fmt.Sprintf("%q", event), nil, o)
		}
	}()
}
