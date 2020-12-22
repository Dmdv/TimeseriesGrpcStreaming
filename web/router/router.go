package router

import (
	"github.com/dmdv/timeseriesgrpcstreaming/web/middleware"
	"github.com/gin-gonic/gin"
	ginmw "github.com/wajox/gin-ext-lib/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		ginmw.Recovery(),
		ginmw.Logger(),
		middleware.NewCORSMiddleware(),
	)

	return r
}
