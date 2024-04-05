package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Uin         string `gorm:"column:uin" `
	Name        string `gorm:"column:name" `
	Password    string `gorm:"column:password"`
	PhoneNumber string `gorm:"column:phone_number" `
	Perm        int    `gorm:"column:perm"`
}

func (obj *User) TableName() string {
	return "User"
}
