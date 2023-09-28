package dao

import "Hybbs-API/Tools"

func SaveChatListDao(sendId int64, receId int64, message string, date string) error {
	tx := Tools.DB.Exec("INSERT INTO chat(c_send_id,c_rece_id,c_message,c_date)VALUES (?,?,?,?)", sendId, receId, message, date)
	return tx.Error
}
