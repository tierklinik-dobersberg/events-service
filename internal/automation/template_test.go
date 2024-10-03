package automation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplateModule(t *testing.T) {
	rt, err := New("test", nil, WithTemplateModule())
	require.NoError(t, err, "creating a new engine should work")

	require.NoError(t, rt.RunScript(`
		const template = require("template")
	`))

	require.NoError(t, rt.RunScript(`
		let t = new template.Template()

		t.register("greet", (s) => {
			console.log("greet called with " + s)
			return "Hello " + s;
		})

		var result = t.exec('{{ "World" | greet }}')

		if (result !== "Hello World") {
			throw new Error("unexpected response: " + result)
		}
	`))

	require.NoError(t, rt.RunScript(`
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
	`))
}
