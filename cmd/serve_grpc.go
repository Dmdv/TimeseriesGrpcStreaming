package cmd

import (
	//"github.com/dmdv/TimeseriesGrpcStreaming/grpc"
	"github.com/spf13/cobra"
)

var ServeGRPCCmd = &cobra.Command{
	Use:     "serve-grpc",
	Aliases: []string{"grpc"},
	Short:   "Start grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		//grpc.Start()
	},
}
