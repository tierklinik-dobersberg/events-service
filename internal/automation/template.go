package automation

import (
	"bytes"
	"fmt"
	"log/slog"
	"text/template"

	"github.com/dop251/goja"
)

type templateInstance struct {
	fn map[string]goja.Callable
}

func WithTemplateModule() EngineOption {
	return func(e *Engine) {
		e.RegisterNativeModuleHelper("template", map[string]any{
			"new": func(runtime *goja.Runtime, call goja.ConstructorCall) *goja.Object {
				slog.Info("template.new", "rt", fmt.Sprintf("%T", runtime), "call", fmt.Sprintf("%T", call))

				instance := &templateInstance{
					fn: make(map[string]goja.Callable),
				}

				call.This.Set("register", func(name string, call goja.Callable) error {
					instance.fn[name] = call
					return nil
				})

				call.This.Set("exec", func(t string, args any) (string, error) {
					tmp, err := template.New(e.name).Parse(t)
					if err != nil {
						return "", err
					}

					fnMap := template.FuncMap{}
					for key, val := range instance.fn {
						fnMap[key] = func(args ...any) any {
							gojaArgs := make([]goja.Value, len(args))
							for i, a := range args {
								gojaArgs[i] = runtime.ToValue(a)
							}

							res, err := val(call.This, gojaArgs...)
							if err != nil {
								panic(err.Error())
							}

							return res.Export()
						}
					}

					var buf = new(bytes.Buffer)
					if err := tmp.Funcs(fnMap).Execute(buf, args); err != nil {
						return "", err
					}

					return buf.String(), nil
				})

				return nil
			},
		})
	}
}
