package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Zones []map[string]interface{} `json:"zones"`
	Posts []map[string]interface{} `json:"posts"`
	Users []map[string]interface{} `json:"users"`
}

func SearchController(ctx *gin.Context) {
	searchMsg := ctx.Query("searchContent")

	zones := dao.SearchZoneDao(searchMsg)
	posts := dao.SearchPostDao(searchMsg)
	users := dao.SearchUserDao(searchMsg)
	// 结果集
	result := Result{Zones: zones, Posts: posts, Users: users}
	ctx.JSON(http.StatusOK, Pojo.Res{Status: "200", Data: result, Msg: "成功"})
}
