package Pojo

type Chat struct {
	CId      int    `gorm:"column:c_id" json:"c_id"`
	CSendId  int    `gorm:"column:c_send_id" json:"sendId"`
	CReceId  int    `gorm:"column:c_rece_id" json:"receId"`
	CMessage string `gorm:"column:c_message" json:"message"`
	Cdate    string `gorm:"column:c_date" json:"date"`
}

func (Chat) TableName() string {
	return "chat"
}
