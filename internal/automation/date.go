package automation

import (
	"time"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/apis/pkg/timeutil"
)

func WithDateModule() EngineOption {
	return func(e *Engine) {
		m := make(map[string]any, len(timeutil.TimeFuncs)+2)

		var rt *goja.Runtime

		e.RunAndBlock(func(r *goja.Runtime) error {
			rt = r

			return nil
		})

		for key, val := range timeutil.TimeFuncs {
			m[key] = func(t time.Time) (goja.Value, error) {
				result := val(t)

				return rt.RunScript(key, "new Date('"+result.Format(time.RFC3339)+"')")
			}
		}

		m["parse"] = func(s string) (time.Time, error) {
			return time.Parse(time.RFC3339, s)
		}
		m["parseStart"] = timeutil.ParseStart
		m["parseEnd"] = timeutil.ParseEnd

		e.RegisterNativeModuleHelper("timeutil", m)
	}
}
