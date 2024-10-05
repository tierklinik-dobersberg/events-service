package fs

import (
	"io/fs"
	"os"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

// Filesystem module
type Instance struct {
	Root fs.FS

	rt *goja.Runtime
}

type Module struct {
	Root fs.FS
}

func (*Module) Name() string { return "fs" }

func (m *Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	root := m.Root

	if root == nil {
		root = os.DirFS(vu.PackagePath())
	}

	mod := &Instance{
		Root: root,
		rt:   vu.Runtime(),
	}

	obj := vu.Runtime().NewObject()

	obj.Set("readFile", mod.readFile)
	obj.Set("readDir", mod.readDir)

	return obj, nil
}

func (m *Instance) readFile(path string) string {
	res, err := fs.ReadFile(m.Root, path)
	if err != nil {
		common.Throw(m.rt, err)
	}

	return string(res)
}

func (m *Instance) readDir(path string) []fs.DirEntry {
	res, err := fs.ReadDir(m.Root, path)
	if err != nil {
		common.Throw(m.rt, err)
	}

	return res
}
