package model

type ReplyANDUser struct {
	RId        int    `gorm:"column:r_id" json:"r_id"`
	UId        int    `gorm:"column:u_id" json:"u_id"`
	PId        int    `gorm:"column:p_id" json:"p_id"`
	RContent   string `gorm:"column:r_content" json:"r_content"`
	RDate      string `gorm:"column:r_date" json:"r_date"`
	RReply     int    `gorm:"column:r_reply" json:"r_reply"`
	U_email    string `gorm:"column:u_email" json:"u_email"`
	U_password string `gorm:"column:u_password" json:"u_password"`
	U_nick     string `gorm:"column:u_nick" json:"u_nick"`
	U_lv       int    `gorm:"column:u_lv" json:"u_lv"`
}
