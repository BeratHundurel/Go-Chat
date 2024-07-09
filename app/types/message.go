package types

import "time"

type Message struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Content    string `json:"content" gorm:"not null"`
	SenderId   int    `json:"senderId" gorm:"not null"`
	ReceiverId int    `json:"receiverId" gorm:"not null"`
	CreatedAt  time.Time
}
