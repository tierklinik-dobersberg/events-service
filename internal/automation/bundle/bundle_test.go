package bundle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

func TestLoadAndPrepare(t *testing.T) {
	bundle, err := Load("./testdata/test")
	require.NoError(t, err)

	assert.Equal(t, &Bundle{
		Path:          "./testdata/test",
		Main:          "./test.js",
		Version:       "1.0.0",
		License:       "ISC",
		ScriptContent: "console.log(\"foobar\")\n",
	}, bundle)

	err = bundle.Prepare(config.Config{}, nil)
	require.NoError(t, err)

	require.NotNil(t, bundle.Runtime())
}

func TestBundleDiscover(t *testing.T) {
	bundles, err := Discover("./testdata")
	require.NoError(t, err)

	assert.Len(t, bundles, 4)
}
