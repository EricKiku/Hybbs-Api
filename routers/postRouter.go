package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func PostRouter(ginServer *gin.Engine) {

	group := ginServer.Group("/post")

	group.POST("/publishPost", controllers.PublishPostController)
	// 获取所有帖子
	group.GET("/getPost", controllers.GetPostController)
	// 更新最后回复
	group.PUT("/updateLast", controllers.UpdateLastUserLastDateController)
	// 获取帖子
	group.GET("/getPostByPlz", controllers.GetPostByPlzController)
}
