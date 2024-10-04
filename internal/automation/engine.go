package automation

import (
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
)

type Engine struct {
	name     string
	loop     *eventloop.EventLoop
	Registry *require.Registry
	ldr      require.SourceLoader
	core     *CoreModule
	baseDir  string
}

type EngineOption func(*Engine)

func WithSourceLoader(ldr require.SourceLoader) EngineOption {
	return func(e *Engine) {
		e.ldr = ldr
	}
}

func WithBaseDirectory(dir string) EngineOption {
	return func(e *Engine) {
		e.baseDir = dir
	}
}

func New(name string, broker Broker, opts ...EngineOption) (*Engine, error) {
	engine := &Engine{
		name: name,
		ldr:  require.DefaultSourceLoader,
	}

	registry := require.NewRegistry(require.WithLoader(func(path string) ([]byte, error) {
		if engine.baseDir != "" {
			path = filepath.Join(engine.baseDir, path)
		}

		return engine.ldr(path)
	}))

	loop := eventloop.NewEventLoop(eventloop.WithRegistry(registry))

	engine.loop = loop
	engine.Registry = registry

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

func (e *Engine) RegisterNativeModuleHelper(name string, obj map[string]any) {
	e.Registry.RegisterNativeModule(name, func(r *goja.Runtime, o *goja.Object) {
		exports, ok := o.Get("exports").(*goja.Object)
		if !ok {
			panic("failed to get exports object")
		}

		for key, value := range obj {
			exports.Set(key, value)
		}
	})
}
