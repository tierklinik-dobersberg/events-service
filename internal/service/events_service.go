package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	connect "github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/events-service/internal/broker"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventsService struct {
	eventsv1connect.UnimplementedEventServiceHandler

	broker *broker.Broker
	l      *slog.Logger
}

type fakeBidiStream struct {
	*connect.ServerStream[eventsv1.Event]
	*connect.Request[eventsv1.SubscribeOnceRequest]

	iter int
}

func (fb *fakeBidiStream) Receive() (*eventsv1.SubscribeRequest, error) {
	if fb.iter >= len(fb.Msg.TypeUrls) {
		return nil, io.EOF
	}

	val := fb.Request.Msg.TypeUrls[fb.iter]

	fb.iter++

	return &eventsv1.SubscribeRequest{
		Kind: &eventsv1.SubscribeRequest_Subscribe{
			Subscribe: val,
		},
	}, nil
}

func NewEventsService(broker *broker.Broker) (*EventsService, error) {
	return &EventsService{broker: broker, l: slog.Default().WithGroup("service")}, nil
}

func (svc *EventsService) Subscribe(ctx context.Context, stream *connect.BidiStream[eventsv1.SubscribeRequest, eventsv1.Event]) error {
	subscriber := broker.NewSubscriber(stream, svc.broker)
	return subscriber.Handle(ctx)
}

func (svc *EventsService) SubscribeOnce(ctx context.Context, req *connect.Request[eventsv1.SubscribeOnceRequest], stream *connect.ServerStream[eventsv1.Event]) error {
	subscriber := broker.NewSubscriber(&fakeBidiStream{stream, req, 0}, svc.broker)
	return subscriber.Handle(ctx)
}

func (svc *EventsService) Publish(ctx context.Context, req *connect.Request[eventsv1.Event]) (*connect.Response[emptypb.Empty], error) {
	svc.l.Debug("received publish request", slog.Any("typeUrl", req.Msg.Event.TypeUrl))

	if evt := req.Msg.GetEvent(); evt == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("missing event field"))
	}

	if err := svc.broker.Publish(req.Msg); err != nil {
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

		if err := svc.broker.Publish(stream.Msg()); err != nil {
			return nil, err
		}
	}

	if stream.Err() != nil {
		return nil, stream.Err()
	}

	return connect.NewResponse(new(emptypb.Empty)), nil
}

var _ eventsv1connect.EventServiceHandler = (*EventsService)(nil)
