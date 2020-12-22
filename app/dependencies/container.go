package dependencies

import "github.com/dmdv/timeseriesgrpcstreaming/grpc"

type Container struct {
	GRPC            *grpc.GRPCService
}
