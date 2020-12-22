package grpc

import (
	"context"
	"time"

	pb "github.com/dmdv/timeseriesgrpcstreaming/grpc/schema"
	"google.golang.org/grpc"
)

const GrpcTimeout = time.Second * 5

type GRPCService struct {
	Client     pb.MeterStreamerClient
	connection *grpc.ClientConn
	host       string
}

func NewGRPCService(host string) (*GRPCService, error) {
	return &GRPCService{
		host: host,
	}, nil
}

func (g *GRPCService) Init() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, g.host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}

	g.Client = pb.NewMeterStreamerClient(conn)
	g.connection = conn

	return
}

func (g *GRPCService) Close() (err error) {
	if g.connection != nil {
		err = g.connection.Close()
		if err != nil {
			return
		}
	}
	return
}
