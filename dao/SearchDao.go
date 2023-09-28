package dao

import "Hybbs-API/Tools"

func SearchZoneDao(searchMsg string) []map[string]interface{} {
	var zones []map[string]interface{}
	var sql = "SELECT * FROM zone WHERE z_name LIKE " + "'" + "%" + searchMsg + "%" + "'"
	Tools.DB.Raw(sql).Scan(&zones)
	return zones
}

func SearchPostDao(searchMsg string) []map[string]interface{} {
	var posts []map[string]interface{}
	var sql = "SELECT * FROM post,zone WHERE post.z_id = zone.z_id AND (post.p_title LIKE " + "'" + "%" + searchMsg + "%" + "'" + " OR " + "p_content LIKE" + "'" + "%" + searchMsg + "%" + "')"
	Tools.DB.Raw(sql).Scan(&posts)
	return posts
}

func SearchUserDao(searchMsg string) []map[string]interface{} {
	var users []map[string]interface{}
	var sql = "SELECT * FROM user WHERE u_nick LIKE " + "'" + "%" + searchMsg + "%" + "'"
	Tools.DB.Raw(sql).Scan(&users)
	return users
}
