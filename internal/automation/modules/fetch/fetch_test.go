package fetch

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

func TestFSModule(t *testing.T) {
	reg := &modules.Registry{}
	reg.Register(&Module{})

	rt, err := automation.New("test", config.Config{}, nil, automation.WithModulsRegistry(reg))

	require.NoError(t, err)

	done := make(chan any)
	rt.Runtime().Set("done", func(response any) {
		log.Printf("%#v", response)
		done <- nil
	})
	rt.Runtime().Set("error", func(error any) {
		done <- error
	})

	_, err = rt.RunScript(`
		fetch("https://ifconfig.co", {
			method: "GET",
			headers: {
				"Accept": "application/json",
			},
		})
		.then(response => response.text())
		.then(done)
		.catch(error)
	`)

	v := <-done
	require.NoError(t, err)
	require.Nil(t, v)
}
