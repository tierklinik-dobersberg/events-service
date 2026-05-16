package service

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/hashicorp/go-multierror"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/events-service/internal/webhook"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WebhookRegistryInterface interface {
	RegisterWebhook(webhook.Webhook) error
	List() []webhook.Webhook
}

type WebhookService struct {
	eventsv1connect.UnimplementedWebhookServiceHandler

	registry WebhookRegistryInterface
}

func NewWebhookService(registry WebhookRegistryInterface) *WebhookService {
	return &WebhookService{
		registry: registry,
	}
}

func (whs *WebhookService) RegisterWebhook(ctx context.Context, req *connect.Request[eventsv1.RegisterWebhookRequest]) (*connect.Response[emptypb.Empty], error) {
	hooks := make([]webhook.Webhook, len(req.Msg.Webhooks))

	for idx, pb := range req.Msg.Webhooks {
		wh := webhook.Webhook{
			Path:                pb.WebhookPath,
			TTL:                 pb.TimeToLive.AsDuration(),
			ExpectedContentType: pb.ExpectedContentType,
			MaxContentLength:    uint64(pb.MaxContentLength),
			ExpectedHeaders:     pb.ExpectedHttpHeaders,
		}

		if err := wh.Prepare(); err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid webhook %s: %w", wh.Path, err))
		}

		hooks[idx] = wh
	}

	merr := new(multierror.Error)
	for _, h := range hooks {
		if err := whs.registry.RegisterWebhook(h); err != nil {
			merr.Errors = append(merr.Errors, fmt.Errorf("failed to register webhook %s: %w", h.Path, err))
		}
	}

	if err := merr.ErrorOrNil(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (whs *WebhookService) ListWebhooks(ctx context.Context, req *connect.Request[eventsv1.ListWebhooksRequest]) (*connect.Response[eventsv1.ListWebhooksResponse], error) {
	hooks := whs.registry.List()

	res := &eventsv1.ListWebhooksResponse{
		Webhooks: make([]*eventsv1.Webhook, len(hooks)),
	}

	for idx, h := range hooks {
		res.Webhooks[idx] = h.ToProto()
	}

	return connect.NewResponse(res), nil
}
