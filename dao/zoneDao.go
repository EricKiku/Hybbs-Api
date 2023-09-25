package dao

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/Tools"
)

// 创建zone
func CreateDao(name string, introduce string, lord int, createDate string) (error, int) {
	var zone Pojo.Zone
	zone.ZName = name
	zone.ZIntroduce = introduce
	zone.ZZonelord = lord
	zone.ZCreateDate = createDate
	// 开启事务

	resoult := Tools.DB.Create(&zone)
	tx := Tools.DB.Begin()
	var id int
	tx.Raw("SELECT LAST_INSERT_ID()").Scan(&id)
	// 提交事务
	tx.Commit()
	return resoult.Error, id
}

// 返回所有分区
func GetZoneDao() []Pojo.Zone_user {
	var zone_user []Pojo.Zone_user
	Tools.DB.Raw("SELECT * FROM zone,user WHERE zone.z_zonelord = user.u_id").Scan(&zone_user)
	return zone_user
}

// 更新分区图标
func UpdateZoneIconDao(id string, iconPath string) error {
	tx := Tools.DB.Exec("UPDATE zone SET z_icon = ? WHERE z_id = ?", iconPath, id)
	return tx.Error
}

// 获取分区的所有帖子
func GetPostByZoneIdDao(z_id int) []Pojo.Post {
	var posts []Pojo.Post
	Tools.DB.Raw("SELECt * FROM post WHERE z_id = ?", z_id).Scan(&posts)
	return posts
}

func UpdateZoneFollowDao(z_id int, _type string) error {
	var res error
	if _type == "add" {
		tx := Tools.DB.Exec("UPDATE zone SET z_follows = z_follows + 1 WHERE z_id = ?", z_id)
		res = tx.Error
	} else if _type == "less" {
		tx := Tools.DB.Exec("UPDATE zone SET z_follows = z_follows -1 where z_id = ?", z_id)
		res = tx.Error
	}
	return res

}

// 获取分区byz_id
func GetZoneByZIdDao(z_id int) map[string]interface{} {
	var zone map[string]interface{}
	Tools.DB.Raw("SELECT * FROM zone,user WHERE zone.z_zonelord = user.u_id AND z_id = ?", z_id).Scan(&zone)
	return zone
}

// // 获取分区by z_zonelord
func GetZoneByZonelordDao(uId int) []Pojo.Zone {
	var zones []Pojo.Zone
	Tools.DB.Raw("SELECT * FROM zone WHERE z_zonelord = ?", uId).Scan(&zones)
	return zones
}
