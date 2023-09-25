package Pojo

type User struct {
	U_id       int    `gorm:"column:u_id" json:"u_id"`
	U_email    string `gorm:"column:u_email" json:"u_email"`
	U_password string `gorm:"column:u_password" json:"u_password"`
	U_nick     string `gorm:"column:u_nick" json:"u_nick"`
	U_lv       int    `gorm:"column:u_lv" json:"u_lv"`
	UAttZone   string `gorm:"column:u_att_zone" json:"u_att_zone"`
	UFensi     int    `gorm:"column:u_fensi" json:"u_fensi"`
	UAttention string `gorm:"column:u_attention" json:"u_attention"`
}

func (User) TableName() string {
	return "user"
}
