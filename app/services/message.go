package services

import (
	"go-chat/app/db"
	"go-chat/app/helpers"
	"go-chat/app/types"
)

func SaveMessage(message types.Message) {
	// Save message to database
	db.Get().Create(&message)
}

func GetMessagesByFriendId(friendId string) ([]types.Message, error) {
	// Get messages from database
	var messages []types.Message

	friendID, err := helpers.ReturnIdAsIntFromString(friendId)
	if err != nil {
		return nil, err
	}

	db.Get().Where("sender_id = ? OR receiver_id = ?", friendID, friendID).Order("created_at DESC").Find(&messages)
	return messages, nil
}
