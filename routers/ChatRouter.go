package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func ChatRouter(ginServer *gin.Engine) {

	group := ginServer.Group("/chat")

	group.GET("/getChatBySendAndReceId", controllers.GetChatBySendIdAndReceIdController)
}
