package fetch

import (
	"github.com/dop251/goja"
	"github.com/elazarl/goproxy"
	"github.com/olebedev/gojax/fetch"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

type Module struct{}

func (*Module) Name() string { return "fetch" }

func (*Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	if err := fetch.Enable(vu.EventLoop(), goproxy.NewProxyHttpServer()); err != nil {
		return nil, err
	}

	return nil, nil
}
