package timeutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

func TestDateModule(t *testing.T) {
	reg := &modules.Registry{}
	reg.Register(&Module{})

	rt, err := automation.New("test", config.Config{}, nil, automation.WithModulsRegistry(reg))

	require.NoError(t, err)

	_, err = rt.RunScript(`
	
		const timeutil = require("timeutil")

		const now = new Date(2024, 9, 10)

		const result = timeutil.startOfMonth(now);

		if (!(result instanceof Date)) {
			throw new Error("expected a date object, got: " + (typeof result) + " " + result.format("2006-01-02"))
		}

		const end = timeutil.endOfMonth(result)
		if (!(end instanceof Date)) {
			throw new Error("expected a date object, got: " + (typeof end) + " " + end.format("2006-01-02"))
		}

		const range = timeutil.parseDayTimeRange("06:30 - 07:00")
		if (range.start == null || range.end == null) {
			throw new Error("expected start and end time to be parsed " + range);
		}
		
		if (range.start.hour != 6 || range.start.minute != 30) {
			throw new Error("expected start to be parsed correctly")
		}

	`)

	require.NoError(t, err)
}
