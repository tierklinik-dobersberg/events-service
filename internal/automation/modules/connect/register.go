package connect

import "github.com/tierklinik-dobersberg/events-service/internal/automation/modules"

func init() {
	modules.Register(&ConnectModule{})
}
