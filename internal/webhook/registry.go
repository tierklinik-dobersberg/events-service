package webhook

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"sync"
	"time"
)

// Registry manages registered Webhooks.
type Registry struct {
	wg  sync.WaitGroup
	log *slog.Logger

	rw       sync.RWMutex
	webhooks map[string]Webhook
}

// NewRegistry returns a new webhook registry.
func NewRegistry(ctx context.Context, log *slog.Logger) *Registry {
	r := &Registry{
		webhooks: map[string]Webhook{},
		log:      log.With("subsystem", "webhook-registry"),
	}

	r.wg.Add(1)
	go r.cleanup(ctx)

	return r
}

func (r *Registry) List() []Webhook {
	r.rw.RLock()
	defer r.rw.RUnlock()

	result := make([]Webhook, 0, len(r.webhooks))

	for _, w := range r.webhooks {
		result = append(result, w)
	}

	return result
}

// RegisterWebhook registers a new webhook at the registry. If the webhook
// is already registered, it's created time is kept, last-updated is set
// to now and the webhook registration is replaced.
func (r *Registry) RegisterWebhook(definition Webhook) error {
	// validate thwe webhook path
	if err := ValidateWebhookPath(definition.Path); err != nil {
		return fmt.Errorf("invalid webhook path definition: %w", err)
	}

	// check and prepare expected header regular expressions
	if len(definition.ExpectedHeaders) > 0 {
		definition.expectedHeadersParsed = make(map[string]*regexp.Regexp)

		for header, re := range definition.ExpectedHeaders {
			if re == "" {
				definition.expectedHeadersParsed[header] = nil
			} else {
				p, err := regexp.CompilePOSIX(re)
				if err != nil {
					return fmt.Errorf("invalid regular expression in expected_http_headers: %q: %w", header, err)
				}

				definition.expectedHeadersParsed[header] = p
			}
		}
	}

	// actually perform the registration
	r.rw.Lock()
	defer r.rw.Unlock()

	if existing, ok := r.webhooks[definition.Path]; ok {
		definition.CreatedAt = existing.CreatedAt
	} else {
		definition.CreatedAt = time.Now()
	}
	definition.LastUpdate = time.Now()

	r.webhooks[definition.Path] = definition

	return nil
}

func (r *Registry) cleanup(ctx context.Context) {
	defer r.wg.Done()

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}

		r.rw.Lock()

		for path, w := range r.webhooks {
			if w.TTL == 0 {
				continue
			}

			if time.Now().After(w.LastUpdate.Add(w.TTL)) {
				r.log.Info("deleting expired webhook", "path", path)
				delete(r.webhooks, path)
			}
		}

		r.rw.Unlock()
	}
}
