package template

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
)

func TestTemplateModule(t *testing.T) {
	reg := &modules.Registry{}
	reg.Register(&Module{})

	rt, err := automation.New("test", config.Config{}, nil, automation.WithModulsRegistry(reg))

	require.NoError(t, err, "creating a new engine should work")

	_, err = rt.RunScript(`
		const template = require("template")
	`)
	require.NoError(t, err)

	_, err = rt.RunScript(`
		let t = new template.Template()

		t.register("greet", (s) => {
			console.log("greet called with " + s)
			return "Hello " + s;
		})

		var result = t.exec('{{ "World" | greet }}')

		if (result !== "Hello World") {
			throw new Error("unexpected response: " + result)
		}
	`)
	require.NoError(t, err)

	_, err = rt.RunScript(`
		t = new template.Template()

		t.register("greet", (s) => {
			if (typeof s === "string") {
				return "Hello " + s
			}

			if (Array.isArray(s)) {
				return "Hello " + s.join(", ")
			}
		})

		var result = t.exec('{{ "World" | greet }}')
		if (result !== "Hello World") {
			throw new Error("unexpected response: " + result)
		}

		var result = t.exec('{{ . | greet }}', ["World", "Moon"])
		if (result !== "Hello World, Moon") {
			throw new Error("unexpected response: " + result)
		}
	`)
	require.NoError(t, err)
}
