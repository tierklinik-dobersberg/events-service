package automation

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/require"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

type mockBroker struct {
	events        []*eventsv1.Event
	subscriptions map[string]chan *eventsv1.Event
}

func (m *mockBroker) Publish(event *eventsv1.Event) error {
	m.events = append(m.events, event)
	return nil
}

func (m *mockBroker) Subscribe(topic string, msgs chan *eventsv1.Event) {
	if m.subscriptions == nil {
		m.subscriptions = make(map[string]chan *eventsv1.Event)
	}

	m.subscriptions[topic] = msgs
}

func Test_CoreModule(t *testing.T) {
	done := make(chan struct{})

	engine, err := New("", config.Config{}, nil, func(e *Engine) {
		e.Run(func(r *goja.Runtime) (goja.Value, error) {
			r.Set("done", func() {
				close(done)
			})
			r.Set("error", func(msg string) {
				t.Error(msg)
			})
			return nil, nil
		})
	})

	require.NoError(t, err, "creating a new engine should not fail")

	engine.Run(func(r *goja.Runtime) (goja.Value, error) {
		_, err := r.RunString(`
		var i = 0;
		var id = scheulde("* * * * *", () => {
			i++;
			error("running ...")

			if (i === 2) {
				clearSchedule(id)
				done()
			}
		})
		`)
		return nil, err
	})

	<-done
}

func TestPubSub(t *testing.T) {
	b := &mockBroker{}

	rt, err := New("test", config.Config{}, b)
	require.NoError(t, err)

	_, err = rt.RunScript(`
	on("tkd.events.v1.Event", () => {})
	`)

	require.NoError(t, err)

	require.NotEmpty(t, b.subscriptions)
	_, ok := b.subscriptions["tkd.events.v1.Event"]
	require.True(t, ok)

	_, err = rt.RunScript(`publish("tkd.tasks.v1.TaskEvent" ,{})`)
	require.NoError(t, err)

	require.NotEmpty(t, b.events)
	require.Equal(t, b.events[0].Event.TypeUrl, "type.googleapis.com/tkd.tasks.v1.TaskEvent")
}
