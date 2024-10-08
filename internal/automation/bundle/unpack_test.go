package bundle

import (
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fileTree(t *testing.T, path string) fstest.MapFS {
	t.Helper()

	dirFS := os.DirFS(path)

	m := fstest.MapFS{}

	err := fs.WalkDir(dirFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		var mode fs.FileMode
		var content []byte

		if d.IsDir() {
			mode = fs.ModeDir
		} else {
			content, err = fs.ReadFile(dirFS, path)
			require.NoError(t, err)
		}

		m[path] = &fstest.MapFile{
			Mode: mode,
			Data: content,
		}

		return nil
	})

	require.NoError(t, err)

	return m
}

var expected = fstest.MapFS{
	".": &fstest.MapFile{
		Mode: fs.ModeDir,
	},

	"test": &fstest.MapFile{
		Mode: fs.ModeDir,
	},

	"test/node_modules": &fstest.MapFile{
		Mode: fs.ModeDir,
	},

	"test/node_modules/test-lib": &fstest.MapFile{
		Mode: fs.ModeDir,
	},

	"test/node_modules/test-lib/lib.js": &fstest.MapFile{
		Data: ([]byte)(`console.log("lib")
`),
	},

	"test/package.json": &fstest.MapFile{
		Data: ([]byte)("{\n\t\"main\": \"./test.js\",\n\t\"version\": \"1.0.0\",\n\t\"license\": \"ISC\"\n}\n"),
	},

	"test/test.js": &fstest.MapFile{
		Data: ([]byte)(`console.log("foobar")
`),
	},
}

func Test_unpackTar(t *testing.T) {
	base, err := unpackTar(false, "./testdata/test.tar")
	require.NoError(t, err, "expected unpackTar to succeed")

	defer os.RemoveAll(base)

	m := fileTree(t, base)

	assert.Equal(t, expected, m)
}

func Test_unpackTarGz(t *testing.T) {
	base, err := unpackTar(true, "./testdata/test.tar.gz")
	require.NoError(t, err, "expected unpackTar to succeed")

	defer os.RemoveAll(base)

	m := fileTree(t, base)

	assert.Equal(t, expected, m)
}

func Test_unpackZip(t *testing.T) {
	base, err := unpackZip("./testdata/test.zip")
	require.NoError(t, err, "expected unpackZip to succeed")

	defer os.RemoveAll(base)

	m := fileTree(t, base)

	assert.Equal(t, expected, m)
}
