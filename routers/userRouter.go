package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(ginServer *gin.Engine) {
	group := ginServer.Group("/user")

	// 设置昵称
	group.POST("/setnick", controllers.SetNickController)
	// 设置关注分区列表
	group.POST("/attention", controllers.SetAttentionZoneController)
	// 获取关注分区列表
	group.GET("/attention", controllers.GetAttentionZoneController)
	// 获取用户细腻
	group.GET("/getUser", controllers.GetUserController)
	// 关注用户
	group.POST("/attentionUser", controllers.AttentionUserController)
	// 取消关注
	group.POST("/cancelAttention", controllers.CancelAttentionUserController)
	// 获取分区和回复数
	group.GET("/getZoneAndPostCount", controllers.GetUserZoneAndPostCountController)
	// 获取用户好友
	group.GET("/getFriend", controllers.GetUserFriendListController)
}
