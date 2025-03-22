package automation

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/dop251/goja"
	cron "github.com/robfig/cron/v3"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	longrunningv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/wellknown"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules/connect"
	"github.com/tierklinik-dobersberg/longrunning-service/pkg/op"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
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
	slog.Info("automation: new schedule registered", "schedule", schedule, "name", c.engine.name)

	res, err := c.scheduler.AddFunc(schedule, func() {
		slog.Info("triggering automation schedule", "schedule", schedule, "name", c.engine.name)

		c.wrapOperation(callable, "schedule:"+schedule, nil)
	})
	if err != nil {
		return -1, err
	}

	return int(res), nil
}

func (c *CoreModule) wrapOperation(callable goja.Callable, kind string, this any, args ...any) {
	cli, err := wellknown.LongRunningService.Create(context.Background(), c.engine.discoverer)

	if err != nil || cli == nil {
		slog.Error("failed to get longrunning service instance", "error", err)

		c.engine.loop.RunOnLoop(func(r *goja.Runtime) {
			this := r.ToValue(this)
			a := make([]goja.Value, len(args))
			for idx, arg := range args {
				a[idx] = r.ToValue(arg)
			}

			callable(this, a...)
		})
	} else {
		slog.Info("got long running service instance")

		op.Wrap(context.Background(), cli, func(context.Context) (any, error) {

			var (
				result    any
				resultErr error
			)

			slog.Info("scheduling operation on event loop")

			c.engine.loop.RunOnLoop(func(r *goja.Runtime) {
				this := r.ToValue(this)
				a := make([]goja.Value, len(args))
				for idx, arg := range args {
					a[idx] = r.ToValue(arg)
				}

				slog.Info("arguments perpare, running goja callable ...")
				gv, err := callable(this, a...)

				if err == nil {
					result = gv.Export()
				} else {
					resultErr = err
				}
			})

			return result, resultErr
		}, func(req *longrunningv1.RegisterOperationRequest) {
			req.Kind = kind
			req.Owner = "automation"
			req.Description = c.engine.name
		})
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

	slog.Info("automation: script successfully subscribed to event topic", "event", event, "name", c.engine.name)

	go func() {
		defer slog.Info("automation: subscription loop closed", "name", c.engine.name)
		for m := range msgs {
			slog.Info("automation: received event, converting from proto-message", "typeUrl", m.Event.TypeUrl, "name", c.engine.name)

			o, err := connect.ConvertProtoMessage(m, c.engine.resolver)
			if err != nil {
				slog.Error("failed to convert protobuf message", "error", err)
				continue
			}

			slog.Info("running automation for event", "typeUrl", m.Event.TypeUrl, "name", c.engine.name)

			c.wrapOperation(callable, "event:"+event, nil, o)
		}
	}()
}
