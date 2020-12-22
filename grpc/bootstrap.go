package grpc

import (
	"github.com/dmdv/timeseriesgrpcstreaming/app/initializers"
)

func StartUp() {
	initializers.InitConfig()
}

func Stop() {
}
