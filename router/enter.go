package router

import (
	"LibraryManagement/global"
	"github.com/gin-gonic/gin"
)

type RGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		global.Log.Warnln(err.Error())
	}

	apiRouter := router.Group("/api/")

	apiRouterGroupApp := RGroup{
		RouterGroup: apiRouter,
	}

	apiRouterGroupApp.AdminRouter()
	apiRouterGroupApp.UserRouter()
	return router
}
