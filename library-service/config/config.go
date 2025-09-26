package config

import (
	"os"
)

type (
	Config struct {
		GRPC
	}

	GRPC struct {
		Port        string `env:"GRPC_PORT"`
		GatewayPort string `env:"GRPC_GATEWAY_PORT"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	cfg.GatewayPort = os.Getenv("GRPC_GATEWAY_PORT")

	return cfg, nil
}
