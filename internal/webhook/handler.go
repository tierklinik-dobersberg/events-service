package webhook

import (
	"io"
	"log/slog"
	"net/http"

	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

type BrokerInterface interface {
	Publish(*eventsv1.Event) error
}

type RegistryInterface interface {
	List() []Webhook
}

type Handler struct {
	registry RegistryInterface
	broker   BrokerInterface
	log      *slog.Logger
}

func NewHandler(log *slog.Logger, broker BrokerInterface, registry RegistryInterface) *Handler {
	return &Handler{
		log:      log.With("subsystem", "webhook"),
		broker:   broker,
		registry: registry,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// we never ever accept webhooks with more that 10k bytes
	if r.ContentLength > 10*1024 {
		h.log.Warn("rejecting webhook request", "error", "request entity to large")
		http.Error(w, "Request to large", http.StatusRequestEntityTooLarge)
		return
	}

	// fetch the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("failed to read HTTP request body", "error", err)
		return
	}

	h.log.Info("received webhook request", "url", r.URL.String())

	// find matching webhooks and publish the events
	foundMatch := false
	hooks := h.registry.List()
	for _, w := range hooks {
		evt, err := w.MatchRequest(r.Context(), h.log, r, body)
		if err != nil {
			h.log.Error("failed to match request against webhook definition", "webhook", w.Path, "error", err)
			continue
		}

		if evt != nil {
			foundMatch = true

			a, err := anypb.New(evt)
			if err != nil {
				h.log.Error("failed to create google.protobuf.Any message", "webhook", w.Path, "error", err)
				continue
			}

			if err := h.broker.Publish(&eventsv1.Event{
				Event: a,
			}); err != nil {
				h.log.Error("failed to publish webhook event", "webhook", w.Path, "error", err)
				continue
			}
		}
	}

	if !foundMatch {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// compile time check
var _ http.Handler = (*Handler)(nil)
