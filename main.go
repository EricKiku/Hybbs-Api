package main

import (
	"Hybbs-API/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	routers.LoginRouter(ginServer)
	routers.UserRouter(ginServer)
	routers.ZoneRouter(ginServer)
	routers.PostRouter(ginServer)
	routers.ReplyRouter(ginServer)
	routers.WsRouter(ginServer)
	routers.ChatRouter(ginServer)
	routers.SearchRouter(ginServer)
	ginServer.Run(":1234")
}
