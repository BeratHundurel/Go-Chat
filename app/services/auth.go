package services

import (
	"go-chat/app/db"
	"go-chat/app/types"
	"go-chat/app/views/auth"
)

func CheckUsername(username string) bool {
	db := db.Get()
	result := db.Limit(1).Find(&types.User{}, username)
	return result.Error == nil
}

func RegisterUser(form auth.RegisterFormValues) {
	user := types.User{
		Username: form.Username,
		Phone:    form.Phone,
		Password: form.Password,
	}

	db.Get().Create(&user)
}

func FindUserByUsername(username string) *types.User {
	var user types.User
	db.Get().Where("username = ?", username).First(&user)
	return &user
}