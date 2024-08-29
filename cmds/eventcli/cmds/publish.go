package cmds

import (
	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"google.golang.org/protobuf/types/known/anypb"
)

func GetPublishCommand(root *cli.Root) *cobra.Command {
	var server string
	cmd := &cobra.Command{
		Use: "publish",
		Run: func(cmd *cobra.Command, args []string) {
			c := eventsv1connect.NewEventServiceClient(cli.NewInsecureHttp2Client(), server)

			pb, err := anypb.New(&commonv1.DayTime{
				Hour:   8,
				Minute: 30,
			})

			if err != nil {
				logrus.Fatal(err.Error())
			}

			logrus.Infof("publishing %q", pb.TypeUrl)
			_, err = c.Publish(root.Context(), connect.NewRequest(&eventsv1.Event{
				Event: pb,
			}))
			logrus.Infof("published %q", pb.TypeUrl)

			if err != nil {
				logrus.Fatalf("failed to publish: %s", err)
			}
		},
	}

	cmd.Flags().StringVar(&server, "server", "http://localhost:8090", "")

	return cmd
}
