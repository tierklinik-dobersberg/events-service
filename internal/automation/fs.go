package automation

import (
	"io/fs"

	"github.com/dop251/goja"
)

// Filesystem module
type FSModule struct {
	Root fs.FS
}

func WithFileSystemModule(root fs.FS) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("fs", func(r *goja.Runtime, o *goja.Object) {
			export, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("expected export object to exist")
			}

			mod := &FSModule{
				Root: root,
			}

			export.Set("readFile", mod.readFile)
			export.Set("readDir", mod.readDir)
		})
	}
}

func (m *FSModule) readFile(path string) (string, error) {
	res, err := fs.ReadFile(m.Root, path)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (m *FSModule) readDir(path string) ([]fs.DirEntry, error) {
	return fs.ReadDir(m.Root, path)
}
