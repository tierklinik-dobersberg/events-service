package cmds

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
)

func GetAutomationCommand(root *cli.Root) *cobra.Command {
	var (
		nodeModules string
	)

	cmd := &cobra.Command{
		Use:  "run-script",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := root.Context()

			opts := []automation.EngineOption{
				automation.WithBoardModule(ctx, root.Boards()),
				automation.WithCallModule(ctx, root.CallService()),
				automation.WithFetchModule(),
				automation.WithNotifyModule(ctx, idmv1connect.NewNotifyServiceClient(cli.NewInsecureHttp2Client(), root.Config().BaseURLS.Idm)),
				automation.WithRolesModule(ctx, root.Roles()),
				automation.WithRosterModule(ctx, root.Roster()),
				automation.WithTaskModule(ctx, root.Tasks()),
				automation.WithUsersModule(ctx, root.Users()),
				automation.WithVoiceMailModule(ctx, pbx3cxv1connect.NewVoiceMailServiceClient(cli.NewInsecureHttp2Client(), root.Config().BaseURLS.CallService)),
			}

			if nodeModules != "" {
				opts = append(opts, automation.WithSourceLoader(func(path string) ([]byte, error) {
					return os.ReadFile(filepath.Join(nodeModules, path))
				}))
			}

			content, err := os.ReadFile(args[0])
			if err != nil {
				logrus.Fatal(err)
			}

			engine, err := automation.New(args[0], nil, opts...)
			if err != nil {
				logrus.Fatal(err)
			}

			engine.RunScript(string(content))

			engine.Stop()
		},
	}

	cmd.Flags().StringVar(&nodeModules, "node-modules", "", "Path to a folder containing node_modules/")

	return cmd
}
