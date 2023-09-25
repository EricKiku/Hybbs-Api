package dao

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/Tools"
	"fmt"
	"strconv"
)

// 设置昵称
func SetNickDao(email string, nick string) error {
	tx := Tools.DB.Exec("UPDATE user SET u_nick = ? WHERE u_email = ?", nick, email)
	return tx.Error
}

// 设置关注列表
func SetAttentionZoneDao(u_att_zone string, u_id int) error {
	tx := Tools.DB.Exec("UPDATE user SET u_att_zone = ? WHERE u_id = ?", u_att_zone, u_id)
	return tx.Error
}

// 获取关注分区列表
func GetAttentionZoneDao(u_id int) string {
	var attention string
	Tools.DB.Raw("SELECT u_att_zone FROM user WHERE u_id = ?", u_id).Scan(&attention)
	return attention
}

// 获取用户信息
func GetUserDao(u_id int) map[string]interface{} {
	var userMsg map[string]interface{}
	Tools.DB.Raw("SELECt * FROM user WHERE u_id = ?", u_id).Scan(&userMsg)
	return userMsg
}

// 关注用户
func AttentionUserDao(opUId int, uId int) error {
	// 更新操作用户的关注列表‘
	// 先获取该用户的关注列表
	var attentionStr string
	Tools.DB.Raw("SELECT u_attention FROM user WHERE u_id=?", opUId).Scan(&attentionStr)
	var upStr = "," + strconv.Itoa(uId) + attentionStr
	tx := Tools.DB.Exec("UPDATE user SET u_attention = ? WHERE u_id = ?", upStr, opUId)

	// 让对方fensi +1

	tx.Exec("UPDATE user SET u_fensi = u_fensi + 1 WHERE u_id = ?", uId)
	return tx.Error
}

// 取消关注
func CancelAttentionUserDao(opUId int, uId int, attStr string) error {
	// 先让对方粉丝-1
	// 让该用户关注列表更新
	tx := Tools.DB.Exec("UPDATE user SET u_fensi = u_fensi - 1 WHERE u_id = ?", uId)
	tx.Exec("UPDATE user SET u_attention = ? WHERE u_id = ?", attStr, opUId)

	return tx.Error
}

// 获取用户的分区数和回复数
func GetUserZoneAndPostCountDao(uId int) (int, int) {
	var zoneCount int
	Tools.DB.Raw("SELECT count(*) FROM zone WHERE z_zonelord = ?", uId).Scan(&zoneCount)
	var replyCount int
	Tools.DB.Raw("SELECT count(*) FROM post WHERE p_lz = ?", uId).Scan(&replyCount)
	fmt.Println(zoneCount, replyCount)
	return zoneCount, replyCount
}

// 获取用户好友
func GetUserFriendListDao(attStr []interface{}) []Pojo.User {
	var users []Pojo.User
	Tools.DB.Where("u_id IN ?", attStr).Find(&users)
	return users
}
