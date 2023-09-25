package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
)

// 创建分区
func CreateController(context *gin.Context) {
	data, _ := context.GetRawData()
	var zone map[string]interface{}
	json.Unmarshal(data, &zone)
	name := zone["name"].(string)
	introduce := zone["introduce"].(string)
	lord := int(zone["lord"].(float64))
	createDate := zone["createDate"].(string)
	err, id := dao.CreateDao(name, introduce, lord, createDate)
	if err != nil {
		context.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "创建失败"})
	} else {
		// 成功后返回 z_id ，用于上传图片
		context.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: id, Msg: "创建成功"})
	}
}

// 上传分区图标
func UploadPicture(context *gin.Context) {
	// 接收图片
	file, _ := context.FormFile("icon")
	// 获取后缀名
	ext := path.Ext(file.Filename)
	// 接收其他参数，接收formdata中的json字符串参数
	_ = context.PostForm("email")
	id := context.PostForm("id")
	// 使用 zoneId 给 分区图标命名
	dst := "H:/HybbsProject/Hybbs/src/zoneIcon/" + id + ext
	err := context.SaveUploadedFile(file, dst)
	// 把图片路径更新到zone数据库的指定zone中
	err = dao.UpdateZoneIconDao(id, dst)
	if err != nil {
		context.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "设置失败"})
	} else {
		context.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "设置成功"})
	}
}

// 返回所有分区
func GetZoneController(ctx *gin.Context) {
	zones := dao.GetZoneDao()
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: zones})
}

// 获取该分区的所有帖子
func GetPostByZoneController(ctx *gin.Context) {
	// 获取 zoneid
	zoneid, _ := strconv.Atoi(ctx.Query("z_id"))
	posts := dao.GetPostByZoneIdDao(zoneid)

	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: posts})

}

// 更新分区关注数
func UpdateZoneFollowController(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var msg map[string]interface{}
	json.Unmarshal(data, &msg)
	var z_id = int(msg["z_id"].(float64))
	var _type = msg["type"].(string)
	res := dao.UpdateZoneFollowDao(z_id, _type)
	if res == nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Msg: "更新成功"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "更新失败"})
	}
}

// 根据z_id 获取分区
func GetZoneByZIdController(ctx *gin.Context) {
	z_id, _ := strconv.Atoi(ctx.Query("z_id"))
	zone := dao.GetZoneByZIdDao(z_id)
	if zone != nil {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: zone, Msg: "获取成功"})
	} else {
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "获取失败"})
	}

}

// 获取分区by z_zonelord
func GetZoneByZonelordController(ctx *gin.Context) {
	uId, _ := strconv.Atoi(ctx.Query("u_id"))

	zones := dao.GetZoneByZonelordDao(uId)
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: zones, Msg: "成功"})
}
