package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func ReplyRouter(ginServer *gin.Engine) {
	group := ginServer.Group("/reply")

	group.POST("/publish", controllers.PublishReplyController)
	// 获取回复根据p_id
	group.GET("/get", controllers.GetReplyByPIdController)
	// 获取回复根据u_id
	group.GET("/getByUId", controllers.GetReplyByUIdController)
}
