package modules

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dop251/goja"
	"github.com/hashicorp/go-multierror"
)

var (
	ErrModuleRegistered = errors.New("module already registered")
)

type Registry struct {
	lock sync.RWMutex

	modules map[string]Module
}

// Register registers a new module at the registry.
func (reg *Registry) Register(mod Module) error {
	reg.lock.Lock()
	defer reg.lock.Unlock()

	if _, ok := reg.modules[mod.Name()]; ok {
		return ErrModuleRegistered
	}

	if reg.modules == nil {
		reg.modules = make(map[string]Module)
	}

	reg.modules[mod.Name()] = mod

	return nil
}

type Instance struct {
	Name   string
	Module *goja.Object
}

// EnableModules enables all modules registered at the registry at the
// specified VU.
func (reg *Registry) EnableModules(vu VU) ([]Instance, error) {
	merr := new(multierror.Error)

	reg.lock.RLock()
	defer reg.lock.RLock()

	instances := make([]Instance, 0, len(reg.modules))

	for _, mod := range reg.modules {
		instance, err := mod.NewModuleInstance(vu)
		if err != nil {
			merr.Errors = append(merr.Errors, fmt.Errorf("%s: %w", mod.Name(), err))
			continue
		}

		if instance == nil {
			continue
		}

		instances = append(instances, Instance{
			Name:   mod.Name(),
			Module: instance,
		})
	}

	if err := merr.ErrorOrNil(); err != nil {
		return instances, err
	}

	r := vu.Registry()

	for _, instance := range instances {
		r.RegisterNativeModule(instance.Name, func(r *goja.Runtime, o *goja.Object) {
			o.Set("exports", instance.Module)
		})
	}

	return instances, merr.ErrorOrNil()
}

// DefaultRegistry is the default, gobal registry at which all modules
// can register.
var DefaultRegistry = &Registry{}

// Register registers a new module at the default registr.y
func Register(mod Module) error {
	return DefaultRegistry.Register(mod)
}
