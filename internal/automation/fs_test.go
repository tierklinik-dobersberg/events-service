package automation

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/require"
)

func TestFSModule(t *testing.T) {
	rt, err := New("test", nil, WithFileSystemModule(fstest.MapFS{
		"foo.text": &fstest.MapFile{
			Data: []byte("foobar"),
		},
		"foo/bar.txt": &fstest.MapFile{
			Data: []byte("bar"),
		},
	}))

	require.NoError(t, err)

	require.NoError(t, rt.RunScript(`
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
	`))
}
