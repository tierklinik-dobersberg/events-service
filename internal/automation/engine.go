package automation

import (
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

type Engine struct {
	name     string
	loop     *eventloop.EventLoop
	registry *require.Registry
	ldr      require.SourceLoader
	baseDir  string
	cfg      config.Config
	rt       *goja.Runtime

	moduleRegistry *modules.Registry
}

func (e *Engine) Registry() *require.Registry {
	return e.registry
}

func (e *Engine) Config() config.Config {
	return e.cfg
}

func (e *Engine) PackagePath() string {
	return e.baseDir
}

func (e *Engine) Runtime() *goja.Runtime {
	return e.rt
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

func WithModulsRegistry(reg *modules.Registry) EngineOption {
	return func(e *Engine) {
		e.moduleRegistry = reg
	}
}

func WithConsole(printer console.Printer) EngineOption {
	return func(e *Engine) {
		obj := e.Runtime().NewObject()
		exports := e.Runtime().NewObject()

		obj.Set("exports", exports)
		console := console.RequireWithPrinter(printer)

		console(e.Runtime(), obj)

		e.Runtime().Set("console", exports)
	}
}

func New(name string, cfg config.Config, broker Broker, opts ...EngineOption) (*Engine, error) {
	engine := &Engine{
		cfg:            cfg,
		name:           name,
		ldr:            require.DefaultSourceLoader,
		moduleRegistry: modules.DefaultRegistry,
	}

	registry := require.NewRegistry(require.WithLoader(func(path string) ([]byte, error) {
		if engine.baseDir != "" {
			path = filepath.Join(engine.baseDir, path)
		}

		return engine.ldr(path)
	}))

	loop := eventloop.NewEventLoop(eventloop.WithRegistry(registry), eventloop.EnableConsole(false))

	engine.loop = loop
	engine.registry = registry

	core := NewCoreModule(engine, broker)

	loop.Run(func(r *goja.Runtime) {
		engine.rt = r

		r.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		core.Enable(r)
	})

	// start the loop before applying any engine options
	engine.loop.Start()

	// load all modules from the module registry
	if engine.moduleRegistry != nil {
		if _, err := engine.moduleRegistry.EnableModules(engine); err != nil {
			return nil, err
		}
	}

	// Apply any engine options
	for _, opt := range opts {
		opt(engine)
	}

	return engine, nil
}

func (e *Engine) EventLoop() *eventloop.EventLoop {
	return e.loop
}

func (e *Engine) Run(fn func(*goja.Runtime) (goja.Value, error)) (goja.Value, error) {
	errCh := make(chan error, 1)
	valueChan := make(chan goja.Value, 1)

	e.loop.RunOnLoop(func(r *goja.Runtime) {
		value, err := fn(r)

		valueChan <- value
		errCh <- err
	})

	return <-valueChan, <-errCh
}

func (e *Engine) RunScript(script string) (goja.Value, error) {
	errCh := make(chan error, 1)
	valueChan := make(chan goja.Value, 1)

	e.loop.RunOnLoop(func(r *goja.Runtime) {
		value, err := r.RunScript("", script)

		valueChan <- value
		errCh <- err
	})

	return <-valueChan, <-errCh
}

func (e *Engine) Stop() int {
	return e.loop.Stop()
}

// Compile time check
var _ modules.VU = (*Engine)(nil)
