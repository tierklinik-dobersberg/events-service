package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/dop251/goja"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/webhook"
)

type Module struct{}

func (*Module) Name() string { return "webhook" }

func (m *Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	vu.Log().Log(context.Background(), slog.LevelInfo, "creating webhook module")

	if vu.WebhookRegistry() == nil {
		vu.Log().Log(context.Background(), slog.LevelError, "no webhook registry defined")
		return nil, nil
	}

	if vu.Broker() == nil {
		vu.Log().Log(context.Background(), slog.LevelError, "no message broker available")
		return nil, nil
	}

	registry := vu.WebhookRegistry()
	broker := vu.Broker()

	obj := vu.Runtime().NewObject()

	obj.Set("register", func(pattern string, contentType string, callable goja.Callable) any {
		wh := webhook.Webhook{
			Path:                pattern,
			ExpectedContentType: contentType,
		}

		if err := wh.Prepare(); err != nil {
			common.Throw(vu.Runtime(), err)
		}

		if _, err := registry.RegisterWebhook(wh); err != nil {
			common.Throw(vu.Runtime(), err)
		}

		msgs := make(chan *eventsv1.Event)
		broker.Subscribe("tkd.events.v1.WebhookEvent", msgs)

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			defer close(msgs)
			defer broker.UnsubscribeAll(msgs)
			defer registry.RemoveWebhook(wh.Path)

			for {
				select {
				case <-ctx.Done():
					return
				case msg := <-msgs:
					var evt eventsv1.WebhookEvent

					if err := msg.Event.UnmarshalTo(&evt); err != nil {
						vu.Log().Log(context.Background(), slog.LevelError, "failed to unpack webhook event")
						continue
					}

					var body any = string(evt.Content)
					if strings.Contains(evt.ContentType, "application/json") {
						if err := json.Unmarshal(evt.Content, &body); err != nil {
							body = string(evt.Content)
							vu.Log().Log(context.Background(), slog.LevelWarn, "failed to parse JSON response", "error", err.Error())
						}
					}

					vu.EventLoop().RunOnLoop(func(r *goja.Runtime) {
						callable(nil, r.ToValue(body), r.ToValue(&evt))
					})
				}
			}
		}()

		return cancel
	})

	return obj, nil
}

func init() {
	if err := modules.Register(&Module{}); err != nil {
		panic(fmt.Sprintf("failed to register webhook module: %s", err))
	}
}
