package automation

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/dop251/goja"
	cron "github.com/robfig/cron/v3"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
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
		c.engine.RunAndBlock(func(r *goja.Runtime) error {
			_, err := callable(nil)

			return err
		})
	})
	if err != nil {
		return -1, err
	}

	return int(res), nil
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

	msg, err := objToProto(obj, md)
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

	slog.Info("automation: script is subscribing to event topic", "event", event, "name", c.engine.name)

	c.broker.Subscribe(event, msgs)

	slog.Info("automation: script successfully subscribed to event topic", "event", event, "name", c.engine.name)

	go func() {
		defer slog.Info("automation: subscription loop closed", "name", c.engine.name)
		for m := range msgs {
			slog.Info("automation: received event, converting from proto-message", "typeUrl", m.Event.TypeUrl, "name", c.engine.name)
			o, err := convertProtoMessage(m)
			if err != nil {
				slog.Error("failed to convert protobuf message", "error", err)
				continue
			}

			slog.Info("running automation for event", "typeUrl", m.Event.TypeUrl, "name", c.engine.name)

			c.engine.loop.RunOnLoop(func(r *goja.Runtime) {
				callable(nil, r.ToValue(o))
			})
		}
	}()
}
