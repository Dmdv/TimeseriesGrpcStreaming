package grpc

import (
	stream "github.com/dmdv/timeseriesgrpcstreaming/grpc/schema"
	//"github.com/rs/zerolog/log"
	//
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
)

func (s* Server) StreamData(rq *stream.Request, srv stream.MeterStreamer_StreamDataServer) error {
	return nil
}
