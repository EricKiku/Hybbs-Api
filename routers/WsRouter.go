package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func WsRouter(ginServer *gin.Engine) {
	group := ginServer.Group("/ws")

	group.GET("/conn", controllers.WsController)
}
