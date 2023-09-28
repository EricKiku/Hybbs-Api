package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func SearchRouter(ginServer *gin.Engine) {

	var group = ginServer.Group("/search")

	group.GET("/get", controllers.SearchController)
}
