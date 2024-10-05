package fs

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

func TestFSModule(t *testing.T) {
	reg := &modules.Registry{}
	reg.Register(&Module{
		Root: fstest.MapFS{
			"foo.text": &fstest.MapFile{
				Data: []byte("foobar"),
			},
			"foo/bar.txt": &fstest.MapFile{
				Data: []byte("bar"),
			},
		},
	})

	rt, err := automation.New("test", config.Config{}, nil, automation.WithModulsRegistry(reg))

	require.NoError(t, err)

	_, err = rt.RunScript(`
		const fs = require("fs")

		const result = fs.readFile("foo.text")
		if (result !== "foobar") {
			throw new Error("unexpected file content")
		}

		const entries = fs.readDir("foo")
		if (entries.length !== 1) {
			throw new Error("unexpected number of fs entries")
		}

		if (entries[0].name() !== "bar.txt") {
			throw new Error("unexpected result from DirEntry.name()")
		}
	`)

	require.NoError(t, err)
}
