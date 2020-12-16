package grpc

import (
	"encoding/csv"
	stream "github.com/dmdv/timeseriesgrpcstreaming/grpc/schema"
	proto "github.com/golang/protobuf/ptypes"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s* Server) StreamData(_ *stream.Request, srv stream.MeterStreamer_StreamDataServer) error {

	meterData, err := ReadCsv("meterusage.csv")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	for i, a := range meterData {

		timeString := a[0]
		valueString := a[1]

		timestamp, err := time.Parse("2006-02-01 15:04:05", timeString)
		if err != nil {
			log.Warn().Msg("Failed to parse: " + timeString)
		}
		timeProto, err := proto.TimestampProto(timestamp)
		if err != nil {
			log.Warn().Msg("Failed to parse: " + timestamp.String())
		}
		value, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			log.Warn().Msg("Failed to parse: " + valueString)
		}

		err = srv.Send(&stream.Data{
			Sequence:  int32(i),
			Timestamp: timeProto,
			Value:     float32(value),
		})

		if err != nil {
			log.Error().Err(err).Send()
		}
	}

	return srv.Send(&stream.Data{
		Sequence:  0,
		Timestamp: nil,
		Value:     0,
	})
}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
