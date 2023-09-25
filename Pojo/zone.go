package Pojo

type Zone struct {
	ZId           int    `gorm:"column:z_id" json:"z_id"`
	ZName         string `gorm:"column:z_name" json:"z_name"`
	ZIcon         string `gorm:"column:z_icon" json:"z_icon"`
	ZIntroduce    string `gorm:"column:z_introduce" json:"z_introduce"`
	ZFollows      int    `gorm:"column:z_follows" json:"z_follows"`
	ZPosts        int    `gorm:"column:z_posts" json:"z_posts"`
	ZZonelord     int    `gorm:"column:z_zonelord" json:"z_zonelord"`
	CurrentOnline int    `gorm:"column:current_online" json:"current_online"`
	ZCreateDate   string `gorm:"column:z_createDate" json:"z_createDate"`
}

func (Zone) TableName() string {
	return "zone"
}
