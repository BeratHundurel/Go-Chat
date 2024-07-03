package types

import "time"

type Message struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Content    string `json:"content"`
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
	CreatedAt  time.Time
}

