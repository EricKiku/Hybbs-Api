package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 回复帖子
func PublishReplyController(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var reply map[string]interface{}
	json.Unmarshal(data, &reply)
	// 用户ID
	var u_id = int(reply["u_id"].(float64))
	// 帖子ID
	var p_id = int(reply["p_id"].(float64))
	// 内容
	var r_content = reply["r_content"].(string)
	// 时间
	var r_date = reply["r_date"].(string)
	// 回复给哪个回复
	var rId = int(reply["r_id"].(float64))

	err, err1 := dao.PublishReplyDao(u_id, p_id, r_content, r_date, rId)
	// 两个有一个不是nil就是回复失败
	if err != nil || err1 != nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "回复失败"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "回复成功"})
	}
}

// 获取帖子根据 帖子ID p_id
func GetReplyByPIdController(ctx *gin.Context) {
	p_id, _ := strconv.Atoi(ctx.Query("p_id"))
	replys := dao.GetReplyByPIdDao(p_id)
	if len(replys) != 0 {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: replys})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "获取失败"})
	}
}

// 获取用户的回复
func GetReplyByUIdController(ctx *gin.Context) {
	u_id, _ := strconv.Atoi(ctx.Query("u_id"))

	result := dao.GetReplyByUIdDao(u_id)

	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: result, Msg: "查询成功"})
}
