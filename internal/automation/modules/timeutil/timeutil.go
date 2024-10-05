package timeutil

import (
	"time"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/apis/pkg/timeutil"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

type Module struct{}

func (*Module) Name() string { return "timeutil" }

func (*Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	obj := vu.Runtime().NewObject()
	rt := vu.Runtime()

	for key, val := range timeutil.TimeFuncs {
		obj.Set(key, func(t time.Time) (goja.Value, error) {
			result := val(t)

			return rt.RunScript(key, "new Date('"+result.Format(time.RFC3339)+"')")
		})
	}

	obj.Set("parse", func(s string) (time.Time, error) {
		return time.Parse(time.RFC3339, s)
	})

	obj.Set("parseStart", timeutil.ParseStart)
	obj.Set("parseEnd", timeutil.ParseEnd)

	return obj, nil
}
