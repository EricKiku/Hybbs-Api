package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 设置昵称
func SetNickController(context *gin.Context) {
	data, _ := context.GetRawData()

	// email、nick
	var userInfo map[string]interface{}

	json.Unmarshal(data, &userInfo)

	email := userInfo["email"].(string)
	nick := userInfo["nick"].(string)

	err := dao.SetNickDao(email, nick)
	if err != nil {
		context.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "设置失败"})
	} else {
		context.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "设置成功"})
	}
}

// 设置关注分区
func SetAttentionZoneController(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var msgMap map[string]interface{}
	json.Unmarshal(data, &msgMap)
	// 获取前端传输的关注列表
	var attention = msgMap["u_att_zone"].(string)
	// 获取用户id
	var u_id = int(msgMap["u_id"].(float64))

	err := dao.SetAttentionZoneDao(attention, u_id)
	if err == nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "设置成功"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "设置失败"})
	}
}

// 获取关注分区列表
func GetAttentionZoneController(ctx *gin.Context) {
	u_id, _ := strconv.Atoi(ctx.Query("u_id"))

	attention := dao.GetAttentionZoneDao(u_id)
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: attention})
}

// 获取用户信息
func GetUserController(ctx *gin.Context) {
	uId, _ := strconv.Atoi(ctx.Query("u_id"))
	userDao := dao.GetUserDao(uId)
	if userDao != nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: userDao, Msg: "成功"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "失败"})
	}
}

// 关注用户；分为两步：把关注的用户id存到该用户的关注列表，再把关注的用户fensi+1
func AttentionUserController(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	var optionUId = int(msg["op_u_id"].(float64))
	var uId = int(msg["u_id"].(float64))

	err1 := dao.AttentionUserDao(optionUId, uId)
	if err == nil && err1 == nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "成功"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "失败"})
	}

}

// 取消关注
func CancelAttentionUserController(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	var optionUId = int(msg["op_u_id"].(float64))
	var uId = int(msg["u_id"].(float64))
	var attStr = msg["attStr"].(string)
	err = dao.CancelAttentionUserDao(optionUId, uId, attStr)
	if err == nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "成功"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "失败"})
	}
}

// 获取用户的分区数和回复数
func GetUserZoneAndPostCountController(ctx *gin.Context) {
	uId, _ := strconv.Atoi(ctx.Query("u_id"))
	fmt.Println(uId)
	zone, reply := dao.GetUserZoneAndPostCountDao(uId)
	var zandr = make(map[string]int)
	zandr["zone"] = zone
	zandr["post"] = reply
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: zandr, Msg: "成功"})
	//if zandr != (ZAndR{}) {
	//	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: zandr, Msg: "成功"})
	//} else {
	//	ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "失败"})
	//}
}

// 获取好友列表
func GetUserFriendListController(ctx *gin.Context) {
	data := ctx.Query("attStr")
	var attStr []interface{}
	json.Unmarshal([]byte(data), &attStr)
	//fmt.Println(attStr)
	//for _, v := range attStr {
	//	fmt.Println(v)
	//}
	users := dao.GetUserFriendListDao(attStr)
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: users, Msg: "成功"})
}
