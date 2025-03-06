package timeutil

import (
	"time"

	"github.com/dop251/goja"
	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	"github.com/tierklinik-dobersberg/apis/pkg/timeutil"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
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

	obj.Set("parseDayTime", func(input string) *commonv1.DayTime {
		res, err := commonv1.ParseDayTime(input)
		if err != nil {
			common.Throw(rt, err)
		}

		return res
	})

	obj.Set("parseDayTimeRange", func(input string) *commonv1.DayTimeRange {
		res, err := commonv1.ParseDayTimeRange(input)
		if err != nil {
			common.Throw(rt, err)
		}

		return res
	})

	return obj, nil
}
