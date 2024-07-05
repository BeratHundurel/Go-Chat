package types

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Phone    string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
}
