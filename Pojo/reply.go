package Pojo

type Reply struct {
	RId      int    `gorm:"column:r_id" json:"r_id"`
	UId      int    `gorm:"column:u_id" json:"u_id"`
	PId      int    `gorm:"column:p_id" json:"p_id"`
	RContent string `gorm:"column:r_content" json:"r_content"`
	RDate    string `gorm:"column:r_date" json:"r_date"`
	RReply   int    `gorm:"column:r_reply" json:"r_reply"`
}

func (Reply) TableName() string {
	return "reply"
}
