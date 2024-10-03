package automation

import (
	"bytes"
	"text/template"

	"github.com/dop251/goja"
)

type templateInstance struct {
	fn map[string]goja.Callable
}

func WithTemplateModule() EngineOption {
	return func(e *Engine) {
		var rt *goja.Runtime
		e.RunAndBlock(func(r *goja.Runtime) error {
			rt = r
			return nil
		})

		e.RegisterNativeModuleHelper("template", map[string]any{
			"Template": func(call goja.ConstructorCall) *goja.Object {
				instance := &templateInstance{
					fn: make(map[string]goja.Callable),
				}

				call.This.Set("register", func(name string, call goja.Callable) error {
					instance.fn[name] = call
					return nil
				})

				call.This.Set("exec", func(t string, args any) (string, error) {
					fnMap := template.FuncMap{}
					for key, val := range instance.fn {
						fnMap[key] = func(args ...any) any {
							gojaArgs := make([]goja.Value, len(args))
							for i, a := range args {
								gojaArgs[i] = rt.ToValue(a)
							}

							res, err := val(call.This, gojaArgs...)
							if err != nil {
								panic(err.Error())
							}

							return res.Export()
						}
					}

					tmp, err := template.New(e.name).Funcs(fnMap).Parse(t)
					if err != nil {
						return "", err
					}

					var buf = new(bytes.Buffer)
					if err := tmp.Execute(buf, args); err != nil {
						return "", err
					}

					return buf.String(), nil
				})

				return nil
			},
		})
	}
}
