package bundle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadBundle(t *testing.T) {
	bundle, err := Load("./testdata/test")
	require.NoError(t, err)

	assert.Equal(t, &Bundle{
		Path:          "./testdata/test",
		Main:          "./test.js",
		Version:       "1.0.0",
		License:       "ISC",
		ScriptContent: "console.log(\"foobar\")\n",
	}, bundle)
}
