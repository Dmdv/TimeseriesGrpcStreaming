package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{}

func RegisterCommands() {
	RootCmd.AddCommand(ServeGRPCCmd)
}

func Execute() {
	RegisterCommands()

	if err := RootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
