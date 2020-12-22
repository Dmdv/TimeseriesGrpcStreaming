package app

import (
	"github.com/dmdv/timeseriesgrpcstreaming/app/dependencies"
	"github.com/dmdv/timeseriesgrpcstreaming/app/initializers"
)

func BuildApplication() (*Application, error) {

	grpcServiceConfig := initializers.InitializeGRPConfig()
	grpcService, err := initializers.InitializeGRPC(grpcServiceConfig)
	if err != nil {
		return nil, err
	}

	container := &dependencies.Container{
		GRPC: grpcService,
	}

	engine := initializers.InitializeRouter(container)
	httpServerConfig := initializers.InitializeHTTPServerConfig(engine)
	server, err := initializers.InitializeHTTPServer(httpServerConfig)
	if err != nil {
		return nil, err
	}

	application := &Application{
		httpServer: server,
		Container:  container,
	}

	return application, nil
}
