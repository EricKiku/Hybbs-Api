package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/Tools"
	"Hybbs-API/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发送邮箱验证码控制器
func SendYzmController(context *gin.Context) {
	email := context.Query("email")

	// 调用发送邮箱验证方法
	mail := Tools.SendMail(email)
	// 返回验证码
	context.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: mail, Msg: ""})

}

// 查询邮箱是否存在控制器
func FindEmailController(context *gin.Context) {
	email := context.Query("email")
	// 验证该邮箱存在吗？为true代表存在
	emailDao := dao.FindEmailDao(email)
	if emailDao {
		context.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "该邮箱已存在!"})
	} else {
		context.JSON(http.StatusOK, Pojo.Res{Status: "200"})
	}
}

// 注册账号
func RegisterController(context *gin.Context) {
	data, _ := context.GetRawData()
	var userInfo map[string]interface{}
	_ = json.Unmarshal(data, &userInfo)
	email := userInfo["email"].(string)
	password := userInfo["password"].(string)
	err := dao.RegisterDao(email, password)
	if err != nil {
		context.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "注册失败"})
	} else {
		context.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "注册成功"})
	}

}

// 登录
func LoginController(context *gin.Context) {
	data, _ := context.GetRawData()
	var userInfo map[string]interface{}
	json.Unmarshal(data, &userInfo)
	email := userInfo["email"].(string)
	password := userInfo["password"].(string)
	user := dao.LoginDao(email, password)
	if user != (Pojo.User{}) {
		context.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: user})
	} else {
		context.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "邮箱或密码错误"})
	}
}

// 根据id登录
func LoginByIdController(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var userid map[string]interface{}
	json.Unmarshal(data, &userid)
	var id = userid["id"].(string)
	user := dao.LoginByIdDao(id)
	if user != (Pojo.User{}) {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: user})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "登录失败"})
	}

}
