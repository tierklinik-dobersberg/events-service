package service

import (
	"context"
	"fmt"

	connect "github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/events-service/internal/broker"
	"github.com/tierklinik-dobersberg/events-service/internal/subscription"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventsService struct {
	eventsv1connect.UnimplementedEventServiceHandler

	broker *broker.Broker
}

func (svc *EventsService) Subscribe(ctx context.Context, stream subscription.Stream) error {
	sub, err := subscription.NewSubscription(stream)
	if err != nil {
		return err
	}

	svc.broker.AddSubscription(sub)

	if err := sub.Start(ctx); err != nil {
		return err
	}

	sub.Wait()

	return nil
}

func (svc *EventsService) Publish(ctx context.Context, req *connect.Request[eventsv1.Event]) (*connect.Response[emptypb.Empty], error) {
	if evt := req.Msg.GetEvent(); evt == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("missing event field"))
	}

	if err := svc.broker.Publish(req.Msg.Event.TypeUrl, req.Msg.Event.GetValue()); err != nil {
		return nil, err
	}

	return connect.NewResponse(new(emptypb.Empty)), nil
}

func (svc *EventsService) PublishStream(ctx context.Context, stream *connect.ClientStream[eventsv1.Event]) (*connect.Response[emptypb.Empty], error) {
	for stream.Receive() {
		if err := stream.Err(); err != nil {
			return nil, err
		}

		if evt := stream.Msg().GetEvent(); evt == nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("missing event field"))
		}

		if err := svc.broker.Publish(stream.Msg().Event.TypeUrl, stream.Msg().Event.Value); err != nil {
			return nil, err
		}
	}

	if stream.Err() != nil {
		return nil, stream.Err()
	}

	return connect.NewResponse(new(emptypb.Empty)), nil
}

var _ eventsv1connect.EventServiceHandler = (*EventsService)(nil)
