package services

import (
	"go-chat/app/db"
	"go-chat/app/types"
	"strconv"
)

func getById(id int) types.User {
	var user types.User
	db.Get().First(&user, id)
	return user
}

func GetAvailableUsers(phone string) []types.User {
	var users []types.User
	db.Get().Where("phone != ?", phone).Find(&users)
	return users
}

func GetUserByPhone(phone string) types.User {
	var user types.User
	db.Get().Where("phone = ?", phone).First(&user)
	return user
}

func AddFriend(user types.User, friendId string) error {
	friendID, err := strconv.Atoi(friendId)
	if err != nil {
		return err
	}
	friend := getById(friendID)

	db.Get().Model(&user).Association("Friends").Append(friend)
	return nil
}
