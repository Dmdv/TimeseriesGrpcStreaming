package app

import (
	"context"
	"github.com/dmdv/timeseriesgrpcstreaming/app/dependencies"
	"github.com/dmdv/timeseriesgrpcstreaming/app/initializers"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Application struct {
	httpServer *http.Server
	Container  *dependencies.Container
}

func InitializeApplication() (*Application, error) {

	initializers.InitializeEnvs()

	if err := initializers.InitializeLogs(); err != nil {
		return nil, err
	}

	return BuildApplication()
}

func (a *Application) Start(cli bool) {
	if !cli {
		if err := a.Container.GRPC.Init(); err != nil {
			log.Panic().Err(err).Msg("GRPC client can not connect to server")
		}
		a.StartHTTPServer()
	}
}

func (a *Application) StartHTTPServer() {
	go func() {
		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("HTTP Server stopped")
		}
	}()
}

func (a *Application) Stop() (err error) {

	if err = a.Container.GRPC.Close(); err != nil {
		return
	}

	return a.httpServer.Shutdown(context.TODO())
}
