package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	MqttURL            string   `env:"MQTT_URL,required"`
	ListenAddress      string   `env:"LISTEN, default=:8090"`
	AdminListenAddress string   `env:"ADMIN_LISTEN, default=:8091"`
	AllowedOrigins     []string `env:"ALLOWED_ORIGINS, default=*"`
	IdmURL             string   `env:"IDM_URL"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
