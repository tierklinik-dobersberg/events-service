package cmds

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/apis/pkg/events"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func GetSubscribeCommand(root *cli.Root) *cobra.Command {
	var server string
	cmd := &cobra.Command{
		Use:  "subscribe [typeurl....]",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ec := events.NewClient(func(ctx context.Context) (eventsv1connect.EventServiceClient, error) {
				return eventsv1connect.NewEventServiceClient(http.DefaultClient, server), nil
			})

			if err := ec.Start(root.Context()); err != nil {
				logrus.Fatalf(err.Error())
			}

			ch, err := ec.Subscribe(root.Context(), protoreflect.FullName(args[0]))
			if err != nil {
				logrus.Fatalf(err.Error())
			}

			for msg := range ch {
				root.Print(msg)
			}
		},
	}

	cmd.Flags().StringVar(&server, "server", "http://localhost:8090", "")

	return cmd
}
