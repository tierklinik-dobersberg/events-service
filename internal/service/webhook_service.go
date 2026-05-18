package service

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/events-service/internal/webhook"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WebhookRegistryInterface interface {
	RegisterWebhook(webhook.Webhook) (webhook.Webhook, error)
	RemoveWebhook(string) bool
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

func (whs *WebhookService) RegisterWebhook(ctx context.Context, req *connect.Request[eventsv1.RegisterWebhookRequest]) (*connect.Response[eventsv1.RegisterWebhookResponse], error) {
	pb := req.Msg.Webhook
	if pb == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("missing webhook definition"))
	}

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

	res, err := whs.registry.RegisterWebhook(wh)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&eventsv1.RegisterWebhookResponse{
		Webhook: res.ToProto(),
	}), nil
}

func (whs *WebhookService) RemoveWebhook(ctx context.Context, req *connect.Request[eventsv1.RemoveWebhookRequest]) (*connect.Response[emptypb.Empty], error) {
	found := whs.registry.RemoveWebhook(req.Msg.WebhookPath)

	if !found {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("webhook with path pattern %s not found", req.Msg.WebhookPath))
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
