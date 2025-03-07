package broker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	connect "github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
)

type SubscriberStream interface {
	Send(*eventsv1.Event) error
	Receive() (*eventsv1.SubscribeRequest, error)
	Peer() connect.Peer
}

type Subscriber struct {
	stream SubscriberStream
	broker *Broker
	log    *slog.Logger
}

func NewSubscriber(stream SubscriberStream, broker *Broker) *Subscriber {
	return &Subscriber{
		stream: stream,
		broker: broker,
		log:    slog.Default().With("subsystem", "subscriber", "peer", stream.Peer().Addr),
	}
}

func (s *Subscriber) Handle(ctx context.Context) error {
	msgs := make(chan *eventsv1.Event, 100)

	go func() {
		// wait for the connection to complete
		<-ctx.Done()

		// unsubscribe from the broker, once returned
		// msgs cannot be used again by the broker and we are
		// safe to close it
		s.broker.UnsubscribeAll(msgs)

		close(msgs)
	}()

	go func() {
		defer s.log.Debug("receive loop finished")

		for {
			msg, err := s.stream.Receive()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					s.log.Error("failed to read message from stream", "error", err.Error())
				}

				return
			}

			switch v := msg.Kind.(type) {
			case *eventsv1.SubscribeRequest_Subscribe:
				s.log.Debug("subscribing to topic", "topic", v.Subscribe)
				s.broker.Subscribe(v.Subscribe, msgs)

			default:
				s.log.Error("unhandled message", "type", fmt.Sprintf("%T", msg.Kind))
			}
		}
	}()

	for m := range msgs {
		if err := s.stream.Send(m); err != nil {
			if !errors.Is(err, io.EOF) {
				s.log.Error("failed to send message over stream", "error", err.Error())
			} else {
				s.log.Info("client disconnected")
			}

			break
		}
	}

	s.log.Debug("subscription completed")

	return nil
}
