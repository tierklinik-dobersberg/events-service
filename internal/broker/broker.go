package broker

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tierklinik-dobersberg/events-service/internal/subscription"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Broker struct {
	connLock sync.Mutex
	conn     mqtt.Client

	subscriptionLock sync.Mutex
	subscriptions    []*subscription.Subscription

	log *slog.Logger
}

func NewBroker(ctx context.Context, u string) (*Broker, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("invalid mqtt URL: %w", err)
	}

	broker := &Broker{
		log: slog.Default().WithGroup("broker"),
	}

	opts := mqtt.NewClientOptions()
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(broker.handleConnect)
	opts.Servers = []*url.URL{
		parsedUrl,
	}

	cli := mqtt.NewClient(opts)

	if token := cli.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return broker, nil
}

func (b *Broker) handleConnect(cli mqtt.Client) {
	b.connLock.Lock()
	defer b.connLock.Unlock()

	b.conn = cli
}

func (b *Broker) resubscribeAll() {
	b.subscriptionLock.Lock()

	topics := make(map[string]struct{})
	for _, s := range b.subscriptions {
		for _, t := range s.Topics() {
			topics[t] = struct{}{}
		}
	}

	b.subscriptionLock.Unlock()

	b.connLock.Lock()
	cli := b.conn
	b.connLock.Unlock()

	for key := range topics {
		cisTopic := makeTopic(key)

		if t := cli.Subscribe(cisTopic, 0, b.handleMessage); t.Wait() && t.Error() != nil {
			b.log.Error("failed to subscribe to topic", slog.Any("error", t.Error().Error()))

			break
		}
	}
}

func (b *Broker) handleMessage(_ mqtt.Client, msg mqtt.Message) {
	if msg.Duplicate() {
		return
	}

	pb := new(anypb.Any)

	if err := proto.Unmarshal(msg.Payload(), pb); err != nil {
		b.log.Error("failed to unmarshal google.protobuf.Any", slog.Any("error", err.Error()))
		return
	}

	b.subscriptionLock.Lock()
	defer b.subscriptionLock.Unlock()

	for _, s := range b.subscriptions {
		s.Publish(pb)
	}
}

func (b *Broker) Publish(typeUrl string, blob []byte) error {
	b.connLock.Lock()
	defer b.connLock.Unlock()

	cisTopic := makeTopic(typeUrl)
	if t := b.conn.Publish(cisTopic, 0, false, blob); t.Wait() && t.Error() != nil {
		return t.Error()
	}

	return nil
}

func (b *Broker) AddSubscription(sub *subscription.Subscription) {
	sub.OnClose(b.handleOnClose)
	sub.OnNewTopic(b.handleNewTopic)

	b.subscriptionLock.Lock()
	defer b.subscriptionLock.Unlock()

	b.subscriptions = append(b.subscriptions, sub)
}

func (b *Broker) handleNewTopic(_ *subscription.Subscription, topic string) {
	b.connLock.Lock()
	defer b.connLock.Unlock()

	b.subscriptionLock.Lock()
	found := false
L:
	for _, s := range b.subscriptions {
		for _, t := range s.Topics() {
			if t == topic {
				found = true
				break L
			}
		}
	}
	b.subscriptionLock.Unlock()

	if found {
		return
	}

	cisTopic := makeTopic(topic)
	if t := b.conn.Subscribe(cisTopic, 0, b.handleMessage); t.Wait() && t.Error() != nil {
		b.log.Error("failed to subscribe to new topic", slog.Any("error", t.Error().Error()), slog.Any("topic", cisTopic))
	}
}

func (b *Broker) handleOnClose(sub *subscription.Subscription) {
	topicCount := make(map[string]int)
	b.subscriptionLock.Lock()
	for idx, s := range b.subscriptions {
		if sub == s {
			b.subscriptions = append(b.subscriptions[:idx], b.subscriptions[idx+1:]...)

			for _, t := range s.Topics() {
				topicCount[t]--
			}
		} else {
			for _, t := range s.Topics() {
				topicCount[t]++
			}
		}
	}
	b.subscriptionLock.Unlock()

	b.connLock.Lock()
	defer b.connLock.Unlock()

	for t, c := range topicCount {
		if c <= 0 {
			if token := b.conn.Unsubscribe(makeTopic(t)); token.Wait() && token.Error() != nil {
				b.log.Error("failed to unsubscribe from topic", slog.Any("error", token.Error().Error()), slog.Any("topic", t))
			}
		}
	}
}

func makeTopic(typeUrl string) string {
	return fmt.Sprintf("cis/protobuf/events/%s", typeUrl)
}
