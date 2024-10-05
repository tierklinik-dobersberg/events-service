package template

import (
	"bytes"
	"text/template"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

type templateInstance struct {
	fn map[string]goja.Callable
}

type Module struct{}

func (*Module) Name() string { return "template" }

func (*Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	obj := vu.Runtime().NewObject()

	obj.Set("Template", func(call goja.ConstructorCall) *goja.Object {
		rt := vu.Runtime()

		instance := &templateInstance{
			fn: make(map[string]goja.Callable),
		}

		call.This.Set("register", func(name string, call goja.Callable) error {
			instance.fn[name] = call
			return nil
		})

		call.This.Set("exec", func(t string, args any) string {
			fnMap := template.FuncMap{}
			for key, val := range instance.fn {
				fnMap[key] = func(args ...any) any {
					gojaArgs := make([]goja.Value, len(args))
					for i, a := range args {
						gojaArgs[i] = rt.ToValue(a)
					}

					res, err := val(call.This, gojaArgs...)
					if err != nil {
						common.Throw(rt, err)
					}

					return res.Export()
				}
			}

			tmp, err := template.New(vu.PackagePath()).Funcs(fnMap).Parse(t)
			if err != nil {
				common.Throw(rt, err)
			}

			var buf = new(bytes.Buffer)
			if err := tmp.Execute(buf, args); err != nil {
				common.Throw(rt, err)
			}

			return buf.String()
		})

		return nil
	})

	return obj, nil
}
