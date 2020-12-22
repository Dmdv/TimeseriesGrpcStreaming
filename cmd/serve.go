package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmdv/timeseriesgrpcstreaming/app"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting")

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		application, err := app.InitializeApplication()

		if err != nil {
			log.Error().Err(err).Msg("can not initialize application")
			return
		}

		application.Start(false)
		log.Info().Msg("Started")
		<-sigchan

		log.Error().Err(application.Stop()).Msg("stop application")

		time.Sleep(time.Second * 5)
		log.Info().Msg("Finished")
	},
}