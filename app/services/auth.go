package services

import (
	"go-chat/app/db"
	"go-chat/app/types"
)

func CheckUsername(username string) bool {
	db := db.Get()
	result := db.Limit(1).Find(&types.User{}, username)
	return result.RowsAffected == 1
}

func RegisterUser(user types.User) {
	db.Get().Create(&user)
}
