package automation

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/require"
)

func Test_CoreModule(t *testing.T) {
	done := make(chan struct{})

	engine, err := New("", nil, func(e *Engine) {
		e.RunAndBlock(func(r *goja.Runtime) error {
			r.Set("done", func() {
				close(done)
			})
			r.Set("error", func(msg string) {
				t.Error(msg)
			})
			return nil
		})
	})

	require.NoError(t, err, "creating a new engine should not fail")

	engine.RunAndBlock(func(r *goja.Runtime) error {
		_, err := r.RunString(`
		var i = 0;
		var id = scheulde("* * * * *", () => {
			i++;
			error("running ...")

			if (i === 2) {
				clearSchedule(id)
				done()
			}
		})
		`)
		return err
	})

	<-done
}
