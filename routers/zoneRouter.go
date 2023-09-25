package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func ZoneRouter(ginServer *gin.Engine) {
	group := ginServer.Group("/zone")

	// 创建一个分区
	group.POST("/create", controllers.CreateController)
	// 上传分区图片
	group.POST("/uploadpicture", controllers.UploadPicture)
	// 返回所有分区
	group.GET("/getzone", controllers.GetZoneController)
	// 获取分区的帖子
	group.GET("/getPostByZone", controllers.GetPostByZoneController)
	// 更新分区关注数
	group.PUT("/updateFollow", controllers.UpdateZoneFollowController)
	// 获取分区byz_id
	group.GET("/getByZId", controllers.GetZoneByZIdController)
	// 获取分区byzonelord
	group.GET("/getZoneByLord", controllers.GetZoneByZonelordController)
}
