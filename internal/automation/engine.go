package automation

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/tierklinik-dobersberg/events-service/internal/broker"
)

type Engine struct {
	name     string
	loop     *eventloop.EventLoop
	Registry *require.Registry
	core     *CoreModule
}

type EngineOption func(*Engine)

func New(name string, broker *broker.Broker, opts ...EngineOption) (*Engine, error) {
	registry := require.NewRegistry(require.WithLoader(func(path string) ([]byte, error) {
		return require.DefaultSourceLoader(path)
	}))

	loop := eventloop.NewEventLoop(eventloop.WithRegistry(registry))

	engine := &Engine{
		loop:     loop,
		Registry: registry,
	}

	core := NewCoreModule(engine, broker)
	engine.core = core

	loop.Run(func(r *goja.Runtime) {
		r.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		core.Enable(r)
	})

	// start the loop before applying any engine options
	engine.loop.Start()

	for _, opt := range opts {
		opt(engine)
	}

	return engine, nil
}

// RunAndBlock schedules a function to execute on the loop and waits
// until execution finished.
func (e *Engine) RunAndBlock(fn func(r *goja.Runtime) error) error {
	err := make(chan error, 1)

	e.loop.RunOnLoop(func(r *goja.Runtime) {
		err <- fn(r)
	})

	return <-err
}

// RunScript schedules a script to execute on the loop.
func (e *Engine) RunScript(script string) error {
	err := make(chan error, 1)

	e.loop.RunOnLoop(func(r *goja.Runtime) {
		_, e := r.RunString(script)

		err <- e
	})

	return <-err
}

func (e *Engine) Stop() int {
	return e.loop.Stop()
}
