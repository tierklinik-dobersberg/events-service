package subscription

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"sync"

	"github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	ErrStarted = errors.New("subscription already started")
)

type Stream = connect.BidiStream[eventsv1.SubscribeRequest, eventsv1.Event]

type OnClosedCallback func(s *Subscription)

type Subscription struct {
	l             sync.Mutex
	protoMessages []string
	stream        Stream
	started       chan struct{}
	closed        chan struct{}
	wg            sync.WaitGroup
	onClose       OnClosedCallback
	sendQueue     chan *eventsv1.Event

	log *slog.Logger
}

func (s *Subscription) Peer() connect.Peer {
	return s.stream.Conn().Peer()
}

func NewSubscription(stream Stream) (*Subscription, error) {
	return &Subscription{
		stream:    stream,
		started:   make(chan struct{}),
		closed:    make(chan struct{}),
		sendQueue: make(chan *eventsv1.Event, 100),
		log:       slog.Default().WithGroup(stream.Peer().Addr),
	}, nil
}

func (s *Subscription) OnClose(fn OnClosedCallback) {
	s.l.Lock()
	defer s.l.Unlock()

	if s.IsClosed() {
		fn(s)

		return
	}

	if s.onClose != nil {
		old := s.onClose
		s.onClose = func(s *Subscription) {
			old(s)
			fn(s)
		}
	} else {
		s.onClose = fn
	}
}

func (s *Subscription) Start(ctx context.Context) error {
	s.l.Lock()
	defer s.l.Unlock()

	select {
	case _, ok := <-s.started:
		if !ok {
			return ErrStarted
		}
	default:
	}

	// ensure we won't get started twice.
	close(s.started)

	s.wg.Add(2)
	go s.handleReceive(ctx)
	go s.handleSend(ctx)

	return nil
}

func (s *Subscription) IsClosed() bool {
	select {
	case _, ok := <-s.closed:
		if !ok {
			return true
		}
	default:
	}

	return false
}

func (s *Subscription) Publish(event *anypb.Any) {
	s.l.Lock()
	defer s.l.Unlock()

	if s.IsClosed() {
		return
	}

	found := false
	for _, m := range s.protoMessages {
		if event.TypeUrl == m {
			found = true
			break
		}
	}

	if !found {
		return
	}

	s.sendQueue <- &eventsv1.Event{
		Event: proto.Clone(event).(*anypb.Any),
	}
}

func (s *Subscription) handleReceive(ctx context.Context) {
	defer s.wg.Done()

	for {
		if ctx.Err() != nil {
			return
		}

		msg, err := s.stream.Receive()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				s.log.Error("failed to send message", slog.Any("error", err.Error()))
			}

			return
		}

		s.handleSubscription(msg)
	}
}

func (s *Subscription) handleSubscription(msg *eventsv1.SubscribeRequest) {
	s.l.Lock()
	defer s.l.Unlock()

	switch v := msg.Kind.(type) {
	case *eventsv1.SubscribeRequest_Subscribe:
		s.protoMessages = append(s.protoMessages, v.Subscribe)

	case *eventsv1.SubscribeRequest_Unsubscribe:
		for idx, m := range s.protoMessages {
			if m == v.Unsubscribe {
				s.protoMessages = append(s.protoMessages[:idx], s.protoMessages[idx+1:]...)
			}
		}

	default:
		s.log.Warn("invalid or unspecified subscription request: %T", v)
	}
}

func (s *Subscription) handleSend(ctx context.Context) {
	defer s.wg.Done()
	defer close(s.closed)

L:
	for {
		select {
		case <-ctx.Done():
			break L

		case msg, ok := <-s.sendQueue:
			if !ok {
				return
			}

			if err := s.stream.Send(msg); err != nil {
				if !errors.Is(err, io.EOF) {
					s.log.Error("failed to send message", slog.Any("error", err.Error()))
				}

				return
			}
		}
	}
}

func (s *Subscription) Wait() {
	s.wg.Wait()
}
