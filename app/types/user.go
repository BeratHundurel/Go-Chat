package types

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Username string  `gorm:"column:username;unique;not null"`
	Phone    string  `gorm:"column:phone;unique;not null"`
	Password string  `gorm:"column:password;not null"`
	Friends  []*User `gorm:"many2many:user_friends"`
}
