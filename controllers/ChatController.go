package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetChatBySendIdAndReceIdController(ctx *gin.Context) {
	sendId, _ := strconv.Atoi(ctx.Query("sendId"))
	receId, _ := strconv.Atoi(ctx.Query("receId"))
	chats := dao.GetChatBySendIdAndReceIdDao(sendId, receId)

	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: chats, Msg: "获取成功"})
}
