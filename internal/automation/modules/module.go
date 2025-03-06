package modules

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
	"github.com/tierklinik-dobersberg/pbtype-server/pkg/protoresolve"
)

type AutomationAnnotation struct {
	ConnectHeaders map[string]string `json:"connectHeaders"`
}

type VU interface {
	// Runtime returns a reference to the actual goja runtime.
	Runtime() *goja.Runtime

	// Registry returns the require registry for the goja runtime.
	Registry() *require.Registry

	// Config returns the configuration object.
	Config() config.Config

	// RunScript schedules a javascript script to be executed on the event loop.
	// It blocks until the script has been executed.
	RunScript(string) (goja.Value, error)

	// Run schedules a callback function to be executed on the event loop.
	// It blocks until the function has been executed.
	Run(func(*goja.Runtime) (goja.Value, error)) (goja.Value, error)

	// PackagePath returns the path of the parent directory that contains
	// the loaded and executed automation package.
	PackagePath() string

	AutomationConfig() AutomationAnnotation

	// EventLoop returns the underlying event loop
	EventLoop() *eventloop.EventLoop

	// Discoverer returns the configured service discoverer
	// or a NO-OP if unconfigured.
	Discoverer() discovery.Discoverer

	// TypeResolver returns the protobuf type resolver which is either backed
	// by protoregistry or by using a pbtype-server instance.
	TypeResolver() protoresolve.Resolver
}

type Module interface {
	// Name returns the name of the module. The name is used
	// to load the module in the javascript environment using
	// require()
	Name() string

	// NewModuleInstance should return a new module instance
	// that is made available to JS code using the require().
	// If nil is returned, the module instance is ignored and will
	// not be made available to the JS runtime.
	// If an error is returned, the runtime will refuse to be created.
	// The VU is already prepared and the underlying eventloop has been
	// started.
	NewModuleInstance(VU) (*goja.Object, error)
}
