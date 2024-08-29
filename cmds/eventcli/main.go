package main

// Import all packages that we expect to send events so all protobuf message
// are available in protoregistry.GlobalFiles
import (
	"github.com/sirupsen/logrus"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/comment/v1"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/customer/v1"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/events-service/cmds/eventcli/cmds"
)

func main() {
	root := cli.New("eventcli")

	root.AddCommand(
		cmds.GetSubscribeCommand(root),
		cmds.GetPublishCommand(root),
	)

	if err := root.Execute(); err != nil {
		logrus.Fatal(err.Error())
	}
}
