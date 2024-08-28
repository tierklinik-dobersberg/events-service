package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/server"
	"github.com/tierklinik-dobersberg/events-service/internal/broker"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
	"github.com/tierklinik-dobersberg/events-service/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		slog.Error("failed to load configuration", slog.Any("error", err.Error()))
		os.Exit(-1)
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

	path, handler := eventsv1connect.NewEventServiceHandler(svc)
	serveMux.Handle(path, handler)

	srv, err := server.CreateWithOptions(cfg.ListenAddress, serveMux)
	if err != nil {
		slog.Error("failed to create server", slog.Any("error", err.Error()))
		os.Exit(-1)
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("failed to listen", slog.Any("error", err.Error()))
		os.Exit(-1)
	}
}
