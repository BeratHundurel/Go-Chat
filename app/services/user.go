package services

import (
	"go-chat/app/db"
	"go-chat/app/helpers"
	"go-chat/app/types"
)

func getById(id int) types.User {
	var user types.User
	db.Get().First(&user, id)
	return user
}

func GetAvailableUsers(user types.User) []types.User {
	var available_users []types.User
	var friends []types.User
	var excluded_ids []int
	excluded_ids = append(excluded_ids, int(user.ID))

	db.Get().Model(&user).Association("Friends").Find(&friends)

	for _, friend := range friends {
		excluded_ids = append(excluded_ids, int(friend.ID))
	}

	db.Get().Where("id NOT IN (?)", excluded_ids).Find(&available_users)
	return available_users
}

func GetUserByPhone(phone string) types.User {
	var user types.User
	db.Get().Where("phone = ?", phone).First(&user)
	return user
}

func AddFriend(user types.User, friendId string) error {
	friendID, err := helpers.ReturnIdAsIntFromString(friendId)
	if err != nil {
		return err
	}
	friend := getById(friendID)
	db.Get().Model(&user).Association("Friends").Append(&friend)
	return nil
}

func GetFriendsByUser(user types.User) []types.User {
	var friends []types.User
	db.Get().Model(&user).Association("Friends").Find(&friends)
	return friends
}

func GetLastMessage(user types.User, friend types.User) types.Message {
	var message types.Message
	db.Get().Where("sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?", user.ID, friend.ID, friend.ID, user.ID).Last(&message)
	return message
}
