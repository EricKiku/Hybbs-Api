package dao

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/Tools"
)

func GetChatBySendIdAndReceIdDao(sendId int, receId int) []Pojo.Chat {
	chat := []Pojo.Chat{}
	Tools.DB.Raw("SELECT * FROM chat WHERE (c_send_id=? AND c_rece_id=?)OR (c_send_id=? AND c_rece_id=?)", sendId, receId, receId, sendId).Scan(&chat)
	return chat
}
