package automation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateModule(t *testing.T) {
	rt, err := New("test", nil, WithDateModule())
	require.NoError(t, err)

	require.NoError(t, rt.RunScript(`
	
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
	`))
}
