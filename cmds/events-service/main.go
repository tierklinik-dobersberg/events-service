package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	connect "github.com/bufbuild/connect-go"
	"github.com/bufbuild/protovalidate-go"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1/tasksv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/auth"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/apis/pkg/cors"
	"github.com/tierklinik-dobersberg/apis/pkg/log"
	"github.com/tierklinik-dobersberg/apis/pkg/server"
	"github.com/tierklinik-dobersberg/apis/pkg/validator"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/broker"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
	"github.com/tierklinik-dobersberg/events-service/internal/service"
	"google.golang.org/protobuf/reflect/protoregistry"

	// Import all proto files from tkd/apis
	_ "github.com/tierklinik-dobersberg/apis/proto"
)

var serverContextKey = struct{ S string }{S: "serverContextKey"}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		slog.Error("failed to load configuration", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	protoValidator, err := protovalidate.New()
	if err != nil {
		slog.Error("failed to prepare protovalidate", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	// TODO(ppacher): privacy-interceptor
	interceptors := connect.WithInterceptors(
		log.NewLoggingInterceptor(),
		validator.NewInterceptor(protoValidator),
	)

	if cfg.IdmURL != "" {
		roleClient := idmv1connect.NewRoleServiceClient(http.DefaultClient, cfg.IdmURL)

		authInterceptor := auth.NewAuthAnnotationInterceptor(
			protoregistry.GlobalFiles,
			auth.NewIDMRoleResolver(roleClient),
			func(ctx context.Context, req connect.AnyRequest) (auth.RemoteUser, error) {
				serverKey, _ := ctx.Value(serverContextKey).(string)

				if serverKey == "admin" {
					return auth.RemoteUser{
						ID:          "service-account",
						DisplayName: req.Peer().Addr,
						RoleIDs:     []string{"idm_superuser"}, // FIXME(ppacher): use a dedicated manager role for this
						Admin:       true,
					}, nil
				}

				return auth.RemoteHeaderExtractor(ctx, req)
			},
		)

		interceptors = connect.WithOptions(interceptors, connect.WithInterceptors(authInterceptor))
	}

	corsConfig := cors.Config{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: true,
	}

	b, err := broker.NewMQTTBroker(ctx, cfg.MqttURL)
	if err != nil {
		slog.Error("failed to connect to MQTT broker", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	svc, err := service.NewEventsService(b)
	if err != nil {
		slog.Error("failed to create EventsService", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	serveMux := http.NewServeMux()

	path, handler := eventsv1connect.NewEventServiceHandler(svc, interceptors)
	serveMux.Handle(path, handler)

	loggingHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			slog.Info("handled request", slog.Any("method", r.Method), slog.Any("path", r.URL.Path), slog.Any("duration", time.Since(start).String()))
		})
	}

	wrapWithKey := func(key string, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), serverContextKey, key))

			next.ServeHTTP(w, r)
		})
	}

	// setup automation framework
	if cfg.ScriptPath != "" {
		// Prepare the automation framework
		engineOptions := []automation.EngineOption{
			automation.WithFetchModule(),
		}

		httpCli := cli.NewInsecureHttp2Client()

		if cfg.IdmURL != "" {
			engineOptions = append(engineOptions, automation.WithUsersModule(ctx, idmv1connect.NewUserServiceClient(httpCli, cfg.IdmURL)))
			engineOptions = append(engineOptions, automation.WithRolesModule(ctx, idmv1connect.NewRoleServiceClient(httpCli, cfg.IdmURL)))
		}

		if cfg.RosterURL != "" {
			engineOptions = append(engineOptions,
				automation.WithRolesModule(ctx, idmv1connect.NewRoleServiceClient(httpCli, cfg.RosterURL)),
			)
		}

		if cfg.CallServiceURL != "" {
			engineOptions = append(engineOptions,
				automation.WithCallModule(ctx, pbx3cxv1connect.NewCallServiceClient(httpCli, cfg.CallServiceURL)),
			)
			engineOptions = append(engineOptions,
				automation.WithVoiceMailModule(ctx, pbx3cxv1connect.NewVoiceMailServiceClient(httpCli, cfg.CallServiceURL)),
			)
		}

		if cfg.TaskServiceURL != "" {
			engineOptions = append(engineOptions,
				automation.WithBoardModule(ctx, tasksv1connect.NewBoardServiceClient(httpCli, cfg.TaskServiceURL)),
			)
			engineOptions = append(engineOptions,
				automation.WithTaskModule(ctx, tasksv1connect.NewTaskServiceClient(httpCli, cfg.TaskServiceURL)),
			)
		}

		dirEntries, err := os.ReadDir(cfg.ScriptPath)
		if err != nil {
			slog.Error("failed to read automation scripts", "error", err)
		} else {
			for _, f := range dirEntries {
				if f.IsDir() {
					continue
				}
				content, err := os.ReadFile(f.Name())
				if err != nil {
					slog.Error("failed to read script file", "error", err, "file", f.Name())
					continue
				}

				engine, err := automation.New(f.Name(), b, engineOptions...)
				if err != nil {
					slog.Error("failed to create engine for script file", "error", err, "file", f.Name())
					continue
				}

				if err := engine.RunScript(string(content)); err != nil {
					slog.Error("failed to initialize engine", "error", err)
				}
			}
		}
	}

	// Create the server
	srv, err := server.CreateWithOptions(cfg.ListenAddress, wrapWithKey("public", loggingHandler(serveMux)), server.WithCORS(corsConfig))
	if err != nil {
		slog.Error("failed to setup server", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	adminServer, err := server.CreateWithOptions(cfg.AdminListenAddress, wrapWithKey("admin", loggingHandler(serveMux)), server.WithCORS(corsConfig))
	if err != nil {
		slog.Error("failed to setup admin-server", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	if err := server.Serve(ctx, srv, adminServer); err != nil {
		slog.Error("failed to serve", slog.Any("error", err.Error()))
		os.Exit(-1)
	}
}
