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
	var available_users []types.User
	var user types.User
	var friends []types.User
	var excluded_ids []int

	db.Get().First(&user, "phone = ?", phone)
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
	friendID, err := strconv.Atoi(friendId)
	if err != nil {
		return err
	}
	friend := getById(friendID)
	db.Get().Model(&user).Association("Friends").Append(&friend)
	return nil
}
