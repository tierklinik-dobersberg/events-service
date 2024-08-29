package broker

import mqtt "github.com/eclipse/paho.mqtt.golang"

type blockingClient struct {
	conn mqtt.Client
}

func (bc *blockingClient) Subscribe(topic string, qos byte, handler mqtt.MessageHandler) error {
	t := bc.conn.Subscribe(topic, qos, handler)
	t.Wait()

	return t.Error()
}

func (bc *blockingClient) Publish(topic string, qos byte, payload []byte) error {
	t := bc.conn.Publish(topic, qos, false, payload)
	t.Wait()

	return t.Error()
}

func (bc *blockingClient) Unsubscribe(topic ...string) error {
	t := bc.conn.Unsubscribe(topic...)
	t.Wait()

	return t.Error()
}

var _ BlockingMQTTClient = (*blockingClient)(nil)
