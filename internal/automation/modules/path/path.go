package path

import (
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

type Module struct{}

func (*Module) Name() string { return "path" }

func (*Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	obj := vu.Runtime().NewObject()

	obj.Set("join", filepath.Join)
	obj.Set("ext", filepath.Ext)
	obj.Set("clean", filepath.Clean)
	obj.Set("abs", filepath.Abs)
	obj.Set("isAbs", filepath.IsAbs)
	obj.Set("dir", filepath.Dir)
	obj.Set("glob", filepath.Glob)
	obj.Set("basename", filepath.Base)

	// TODO(ppacher): add walk function

	return obj, nil
}
