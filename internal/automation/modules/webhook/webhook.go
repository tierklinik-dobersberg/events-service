package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
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

	obj := vu.Runtime().NewObject()

	obj.Set("register", func(pattern string, contentType string, callable goja.Callable) any {
		return registerWebhook(vu, pattern, contentType, false, callable)
	})

	obj.Set("registerWithResponse", func(pattern string, contentType string, callable goja.Callable) any {
		return registerWebhook(vu, pattern, contentType, true, callable)
	})

	return obj, nil
}

func registerWebhook(vu modules.VU, pattern string, contentType string, hijack bool, callable goja.Callable) any {
	registry := vu.WebhookRegistry()

	wh := webhook.Webhook{
		Path:                pattern,
		ExpectedContentType: contentType,
	}

	if err := wh.Prepare(); err != nil {
		common.Throw(vu.Runtime(), err)
	}

	handler := func(evt *eventsv1.WebhookEvent, w http.ResponseWriter) chan bool {
		result := make(chan bool)

		var body any = string(evt.Content)
		if strings.Contains(evt.ContentType, "application/json") {
			if err := json.Unmarshal(evt.Content, &body); err != nil {
				body = string(evt.Content)
				vu.Log().Log(context.Background(), slog.LevelWarn, "failed to parse JSON response", "error", err.Error())
			}
		}

		sendResponse := func(code int, contentType string, payload string) {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(code)

			w.Write([]byte(payload))
		}

		vu.EventLoop().RunOnLoop(func(r *goja.Runtime) {
			callable(nil, r.ToValue(body), r.ToValue(&evt), r.ToValue(sendResponse))

			result <- hijack
		})

		return result
	}

	if _, err := registry.RegisterWebhookWithHandler(wh, handler); err != nil {
		common.Throw(vu.Runtime(), err)
	}

	return nil
}

func init() {
	if err := modules.Register(&Module{}); err != nil {
		panic(fmt.Sprintf("failed to register webhook module: %s", err))
	}
}
