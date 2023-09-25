package routers

import (
	"Hybbs-API/controllers"
	"github.com/gin-gonic/gin"
)

func LoginRouter(ginServer *gin.Engine) {
	group := ginServer.Group("/login")

	// 发送验证码
	group.GET("/yzm", controllers.SendYzmController)
	// 验证邮箱存在
	group.GET("/emailExist", controllers.FindEmailController)
	// 注册
	group.POST("/register", controllers.RegisterController)
	// 登录
	group.POST("/login", controllers.LoginController)
	// 根据id登录
	group.POST("/loginById", controllers.LoginByIdController)
}
