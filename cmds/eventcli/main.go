package main

// Import all packages that we expect to send events so all protobuf message
// are available in protoregistry.GlobalFiles
import (
	"github.com/sirupsen/logrus"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	_ "github.com/tierklinik-dobersberg/apis/proto"
	"github.com/tierklinik-dobersberg/events-service/cmds/eventcli/cmds"
)

func main() {
	root := cli.New("eventcli")

	root.AddCommand(
		cmds.GetSubscribeCommand(root),
		cmds.GetPublishCommand(root),
		cmds.GetAutomationCommand(root),
	)

	if err := root.Execute(); err != nil {
		logrus.Fatal(err.Error())
	}
}
