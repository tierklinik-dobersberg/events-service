package automation

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/require"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
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

	engine, err := New("", nil, func(e *Engine) {
		e.RunAndBlock(func(r *goja.Runtime) error {
			r.Set("done", func() {
				close(done)
			})
			r.Set("error", func(msg string) {
				t.Error(msg)
			})
			return nil
		})
	})

	require.NoError(t, err, "creating a new engine should not fail")

	engine.RunAndBlock(func(r *goja.Runtime) error {
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
		return err
	})

	<-done
}

func TestPubSub(t *testing.T) {
	b := &mockBroker{}

	rt, err := New("test", b)
	require.NoError(t, err)

	require.NoError(t, rt.RunScript(`
	
	on("tkd.events.v1.Event", () => {})

	`))

	require.NotEmpty(t, b.subscriptions)
	_, ok := b.subscriptions["tkd.events.v1.Event"]
	require.True(t, ok)

	require.NoError(t, rt.RunScript(`publish("tkd.tasks.v1.TaskEvent" ,{})`))

	require.NotEmpty(t, b.events)
	require.Equal(t, b.events[0].Event.TypeUrl, "type.googleapis.com/tkd.tasks.v1.TaskEvent")
}
