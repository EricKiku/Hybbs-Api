package Pojo

type Post struct {
	PId           int    `gorm:"column:p_id" json:"p_id"`
	ZId           int    `gorm:"column:z_id" json:"z_id"`
	PLz           int    `gorm:"column:p_lz" json:"p_lz"`
	PTitle        string `gorm:"column:p_title" json:"p_title"`
	PContent      string `gorm:"column:p_content" json:"p_content"`
	PDate         string `gorm:"column:p_date" json:"p_date"`
	PReply        int    `gorm:"column:p_reply" json:"p_reply"`
	UName         string `gorm:"column:u_name" json:"u_name"`
	LastUId       int    `gorm:"column:last_u_id" json:"last_u_id"`
	LastUName     string `gorm:"column:last_u_name" json:"last_u_name"`
	LastReplyDate string `gorm:"column:last_reply_date" json:"last_reply_date"`
}

func (Post) TableName() string {
	return "post"
}
