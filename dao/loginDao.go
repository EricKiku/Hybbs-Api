package dao

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/Tools"
)

// 查找有无该邮箱
func FindEmailDao(email string) bool {
	var users []Pojo.User
	Tools.DB.Raw("SELECT u_id from user where u_email = ?", email).Scan(&users)
	if len(users) > 0 {
		return true
	} else {
		return false
	}
}

// 注册
func RegisterDao(email string, password string) error {
	insertRes := Tools.DB.Exec("INSERT INTO user(u_email,u_password) VALUES (?,?)", email, password)
	return insertRes.Error
}

// 登录
func LoginDao(email string, password string) Pojo.User {

	var user Pojo.User

	Tools.DB.Raw("SELECT * FROM user where u_email = ? AND u_password = ?", email, password).Scan(&user)
	return user
}

// 根据id 登录
func LoginByIdDao(id string) Pojo.User {
	var user Pojo.User
	Tools.DB.Raw("SELECT * FROM user WHERE u_id = ?", id).Scan(&user)
	return user
}
