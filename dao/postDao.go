package dao

import (
	"Hybbs-API/Tools"
)

// 发表帖子
func PublishPostDao(z_id int, p_lz int, u_name string, p_title string, p_content string, p_date string) bool {
	// 新增数据
	tx := Tools.DB.Exec("INSERT INTO post(z_id,p_lz,p_title,p_content,p_date,u_name)VALUES (?,?,?,?,?,?)", z_id, p_lz, p_title, p_content, p_date, u_name)
	//更新分区的帖子数+1，根据 z_id
	Tools.DB.Exec("UPDATE zone SET z_posts = z_posts + 1 WHERE z_id = ?", z_id)
	if tx.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

// 获取所有帖子
func GetPostDao() []map[string]interface{} {
	var posts []map[string]interface{}
	Tools.DB.Raw("SELECT * FROM post,zone,user WHERE post.z_id = zone.z_id AND post.p_lz = user.u_id").Scan(&posts)
	return posts
}

// 更新最后回复信息
func UpdateLastUserLastDateDao(last_u_id int, last_u_name string, last_date string, p_id int) error {
	tx := Tools.DB.Exec("UPDATE post SET last_u_id = ?,last_u_name = ?,last_reply_date=? WHERE p_id = ?", last_u_id, last_u_name, last_date, p_id)
	return tx.Error
}

// 获取帖子
func GetPostByPLzDao(pLz int) []map[string]interface{} {
	var posts []map[string]interface{}
	Tools.DB.Raw("SELECT * FROM post,zone WHERE p_lz = ? AND post.z_id = zone.z_id", pLz).Scan(&posts)
	return posts
}
