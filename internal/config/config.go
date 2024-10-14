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

	ScriptPath    string `env:"AUTOMATION_PATH"`
	IdmURL        string `env:"IDM_URL"`
	TypeServerURL string `env:"TYPE_SERVER"`

	// format: <scheme>://<host>:<port>/<fully-qualified-protobuf-service-name>
	ConnectServices []string `env:"SERVICES"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
