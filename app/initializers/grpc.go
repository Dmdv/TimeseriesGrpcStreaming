package initializers

import (
	"github.com/dmdv/timeseriesgrpcstreaming/grpc"
	"github.com/gobuffalo/envy"
)

const (
	GRPCHost = "TIMESERIES__GRPC__SERVER"

	DefaultGRPCHost = "127.0.0.1:50051"
)

type GRPCServiceConfig struct {
	Host string
}

func InitializeGRPConfig() *GRPCServiceConfig {
	return &GRPCServiceConfig{
		Host: envy.Get(GRPCHost, DefaultGRPCHost),
	}
}

func InitializeGRPC(cfg *GRPCServiceConfig) (*grpc.GRPCService, error) {
	return grpc.NewGRPCService(cfg.Host)
}
