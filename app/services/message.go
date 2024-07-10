package services

import (
	"go-chat/app/db"
	"go-chat/app/types"
)

func SaveMessage(message types.Message) {
	// Save message to database
	db.Get().Create(&message)
}

func GetMessagesByFriendId(friendID int) ([]types.Message, error) {
	var messages []types.Message
	db.Get().Where("sender_id = ? OR receiver_id = ?", friendID, friendID).Order("created_at ASC").Find(&messages)
	return messages, nil
}
