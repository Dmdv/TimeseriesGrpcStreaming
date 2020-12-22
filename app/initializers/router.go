package initializers

import (
	"github.com/dmdv/timeseriesgrpcstreaming/app/dependencies"
	"github.com/dmdv/timeseriesgrpcstreaming/web/router"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(container *dependencies.Container) *gin.Engine {
	r := router.NewRouter()

	//ctrls := buildControllers(container)
	//
	//for i := range ctrls {
	//	ctrls[i].DefineRoutes(r)
	//}

	return r
}

//func buildControllers(container *dependencies.Container) []apiv1.Controller {
//	return []apiv1.Controller{
//	}
//}
