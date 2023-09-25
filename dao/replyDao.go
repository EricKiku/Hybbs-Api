package dao

import (
	"Hybbs-API/Tools"
	"Hybbs-API/model"
)

func PublishReplyDao(u_id int, p_id int, r_content string, r_date string, rId int) (error, error) {
	// 插入评论
	tx := Tools.DB.Exec("INSERT INTO reply(u_id,p_id,r_content,r_date,r_reply)VALUES(?,?,?,?,?)", u_id, p_id, r_content, r_date, rId)
	// 给帖子回复数+1
	exec := Tools.DB.Exec("UPDATE post SET p_reply = p_reply + 1 WHERE p_id = ?", p_id)
	return tx.Error, exec.Error
}

// 获取回复
func GetReplyByPIdDao(p_id int) []model.ReplyANDUser {
	var replys []model.ReplyANDUser
	Tools.DB.Raw("SELECT *,user.u_nick FROM reply,user WHERE p_id = ? AND reply.u_id = user.u_id", p_id).Scan(&replys)
	return replys
}

// 获取回复根据 u_id
func GetReplyByUIdDao(u_id int) []map[string]interface{} {
	var result []map[string]interface{}
	Tools.DB.Raw("SELECT * FROM reply,post,zone WHERE reply.u_id = ? AND reply.p_id = post.p_id AND post.z_id = zone.z_id", u_id).Scan(&result)
	return result
}
