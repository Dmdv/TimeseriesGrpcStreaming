package grpc

import (
	"log"
	"net"

	stream "github.com/dmdv/timeseriesgrpcstreaming/grpc/schema"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Server struct {
	stream.UnimplementedStreamerServer
}


func Start() {
	StartUp()

	lis, err := net.Listen("tcp", viper.GetString("grpc.listen"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	stream.RegisterStreamerServer(s, &Server{
	})

	if err := s.Serve(lis); err != nil {
		Stop()
		log.Fatalf("failed to serve: %v", err)
	} else {
		Stop()
	}
}