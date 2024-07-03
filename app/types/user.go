package types

type User struct {
    ID       uint    `gorm:"primaryKey"`
    UserName string `gorm:"column:user_name"`
    Phone    string `gorm:"column:phone"`
    PassWord string `gorm:"column:pass_word"`
}