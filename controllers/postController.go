package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 发表帖子
func PublishPostController(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var postM map[string]interface{}
	_ = json.Unmarshal(data, &postM)
	// 获取 分区id
	var z_id = int(postM["z_id"].(float64))
	// 获取创建者id(楼主)
	var p_lz = int(postM["u_id"].(float64))
	// 获取 昵称
	var u_name = postM["u_name"].(string)
	// 获取帖子title
	var p_title = postM["title"].(string)
	// 获取帖子内容
	var p_content = postM["content"].(string)
	// 获取发表日期
	var p_date = postM["date"].(string)

	bool := dao.PublishPostDao(z_id, p_lz, u_name, p_title, p_content, p_date)
	if bool {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "创建失败"})
	}
}

// 获取所有帖子
func GetPostController(ctx *gin.Context) {
	posts := dao.GetPostDao()

	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: posts})
}

// 更新帖子最后回复人和最后回复时间
func UpdateLastUserLastDateController(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	p_id := int(msg["p_id"].(float64))
	last_u_id := int(msg["u_id"].(float64))
	last_u_name := msg["u_name"].(string)
	last_date := msg["date"].(string)
	err = dao.UpdateLastUserLastDateDao(last_u_id, last_u_name, last_date, p_id)
	if err != nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "更新失败"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "更新成功"})
	}
}

// 获取帖子
func GetPostByPlzController(ctx *gin.Context) {
	pLz, _ := strconv.Atoi(ctx.Query("p_lz"))
	posts := dao.GetPostByPLzDao(pLz)
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: posts, Msg: "成功"})
}
