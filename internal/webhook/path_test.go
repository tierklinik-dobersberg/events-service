package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWebhookPath(t *testing.T) {
	cases := []struct {
		I string            // Input
		P string            // Pattern
		A map[string]string // Path Parameters
		T string            // Trailing
		M bool              // Matches
	}{
		{
			"provet/100/imaging",
			"provet/{id}/imaging",
			map[string]string{
				"id": "100",
			},
			"",
			true,
		},
		{
			"provet/100/imaging/",
			"provet/{id}/imaging",
			map[string]string{
				"id": "100",
			},
			"",
			true,
		},
		{
			"provet/100/imaging",
			"/provet/{id}/imaging/",
			map[string]string{
				"id": "100",
			},
			"",
			true,
		},
		{
			"/provet/100/imaging",
			"provet/{id}/imaging",
			map[string]string{
				"id": "100",
			},
			"",
			true,
		},
		{
			"provet/100/imaging",
			"/provet/{id}/imaging",
			map[string]string{
				"id": "100",
			},
			"",
			true,
		},
		{
			"foo/bar/baz/trailing/path",
			"foo/{bar}/baz/{#}",
			map[string]string{
				"bar": "bar",
			},
			"trailing/path",
			true,
		},
		{
			"foo/bar",
			"foo/baz",
			nil,
			"",
			false,
		},
		{
			"foo/baz",
			"foo/{bar}",
			map[string]string{
				"bar": "baz",
			},
			"",
			true,
		},
		{
			"foo/bar/baz",
			"foo/{bar}",
			nil,
			"",
			false,
		},
		{
			"foo/bar/baz",
			"foo/bar",
			nil,
			"",
			false,
		},
	}

	for _, c := range cases {
		aA, aT, aM := ParseWebhookPath(c.P, c.I)

		assert.Equal(t, c.M, aM, "match result mismatch")
		assert.Equal(t, c.T, aT, "trailing result mismatch")
		assert.Equal(t, c.A, aA, "path parameter mismatch")
	}
}
