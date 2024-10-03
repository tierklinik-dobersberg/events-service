package broker

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"google.golang.org/protobuf/proto"
)

type BlockingMQTTClient interface {
	Subscribe(string, byte, mqtt.MessageHandler) error
	Unsubscribe(...string) error
	Publish(string, byte, bool, []byte) error
}

type Broker struct {
	connLock sync.Mutex
	conn     BlockingMQTTClient

	l         sync.RWMutex
	receivers map[string][]chan *eventsv1.Event
	topics    map[string]struct{}

	retainedMsgs map[string]*eventsv1.Event

	log *slog.Logger
}

func NewMQTTBroker(ctx context.Context, u string) (*Broker, error) {
	broker, err := NewBroker(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create broker: %w", err)
	}

	opts := mqtt.NewClientOptions()
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(broker.HandleOnConnect)
	opts.AddBroker(u)

	cli := mqtt.NewClient(opts)

	if token := cli.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return broker, nil
}

func NewBroker(ctx context.Context, cli BlockingMQTTClient) (*Broker, error) {
	broker := &Broker{
		log:          slog.Default().With("subsystem", "broker"),
		conn:         cli,
		receivers:    make(map[string][]chan *eventsv1.Event),
		topics:       make(map[string]struct{}),
		retainedMsgs: make(map[string]*eventsv1.Event),
	}

	return broker, nil
}

func (b *Broker) HandleOnConnect(cli mqtt.Client) {
	b.connLock.Lock()
	defer b.connLock.Unlock()

	b.log.Info("successfully connected to MQTT")

	b.conn = &blockingClient{cli}

	// reconnect to all topics
	b.l.RLock()
	defer b.l.RUnlock()
	for t := range b.topics {
		topic := makeTopic(t)

		if err := b.conn.Subscribe(topic, 0, b.handleMessage); err != nil {
			b.log.Error("failed to re-subscribe to topic", "topic", topic, "error", err)
		} else {
			b.log.Info("successfully re-subscribed to topic", "topic", topic)
		}
	}
}

func (b *Broker) Subscribe(typeUrl string, msgs chan *eventsv1.Event) {
	b.l.Lock()
	defer b.l.Unlock()

	b.receivers[typeUrl] = append(b.receivers[typeUrl], msgs)

	// immediately send any retained message for that typeUrl
	if msg, ok := b.retainedMsgs[typeUrl]; ok {
		msgs <- msg
	}

	// start to actually subscribe to the topic
	if _, ok := b.topics[typeUrl]; !ok {
		go b.subscribe(typeUrl)
	}
}

func (b *Broker) subscribe(typeUrl string) {
	b.connLock.Lock()
	defer b.connLock.Unlock()

	topic := makeTopic(typeUrl)
	if err := b.conn.Subscribe(topic, 0, b.handleMessage); err != nil {
		b.log.Error("failed to subscribe to topic", "topic", topic)
	} else {
		b.log.Info("successfully subscribed to topic", "topic", topic)
		b.l.Lock()
		defer b.l.Unlock()

		b.topics[typeUrl] = struct{}{}
	}
}

func (b *Broker) UnsubscribeAll(msgs chan *eventsv1.Event) {
	b.l.Lock()
	defer b.l.Unlock()

	var topicCleanup []string
	for key, receivers := range b.receivers {
		for idx, r := range receivers {
			if r == msgs {

				b.receivers[key] = append(receivers[:idx], receivers[idx+1:]...)

				b.log.Info("removing subscriber from topic", "topic", key, "receiverCount", len(b.receivers[key]))
			}

			if len(b.receivers[key]) == 0 {
				b.log.Info("marking topic for cleanup", "topic", key)
				topicCleanup = append(topicCleanup, key)
			}
		}
	}

	if len(topicCleanup) > 0 {
		for _, t := range topicCleanup {
			delete(b.topics, t)
			delete(b.retainedMsgs, t)
		}

		go func() {
			b.connLock.Lock()
			defer b.connLock.Unlock()

			for idx, t := range topicCleanup {
				topicCleanup[idx] = makeTopic(t)
			}

			if err := b.conn.Unsubscribe(topicCleanup...); err != nil {
				b.log.Error("failed to unsubscribe from unused topics", "error", err)
			}

			b.log.Info("successfully unsubscribed from unused topics", "topics", topicCleanup)
		}()
	}
}

func (b *Broker) Publish(evt *eventsv1.Event) error {
	blob, err := proto.Marshal(evt)
	if err != nil {
		return fmt.Errorf("failed to marshal protobuf: %w", err)
	}

	b.connLock.Lock()
	defer b.connLock.Unlock()

	if b.conn == nil {
		return errors.New("not yet connected, please try again later.")
	}

	topic := makeTopic(evt.Event.TypeUrl)

	if err := b.conn.Publish(topic, 0, evt.Retained, blob); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	b.log.Info("published new message", "topic", topic)

	return nil
}

func (b *Broker) handleMessage(_ mqtt.Client, msg mqtt.Message) {
	var pb = new(eventsv1.Event)

	if err := proto.Unmarshal(msg.Payload(), pb); err != nil {
		b.log.Error("failed to unmarshal protobuf message", "error", err.Error())
		return
	}

	typeUrl := strings.TrimPrefix(pb.Event.TypeUrl, "type.googleapis.com/")
	b.log.Info("received new event from mqtt", "typeUrl", typeUrl, "topic", msg.Topic())

	b.l.RLock()
	defer b.l.RUnlock()

	if msg.Retained() {
		pb.Retained = true

		b.retainedMsgs[typeUrl] = pb
	}

	for _, m := range b.receivers[typeUrl] {
		select {
		case m <- proto.Clone(pb).(*eventsv1.Event):
		case <-time.After(time.Second * 5):
			b.log.Info("failed to dispatch event, receiver busy")
		}
	}
}

func makeTopic(typeUrl string) string {
	typeUrl = strings.TrimPrefix(typeUrl, "type.googleapis.com/")

	return fmt.Sprintf("cis/protobuf/events/%s", typeUrl)
}
