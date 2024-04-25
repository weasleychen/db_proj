package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`

	Uin         string `gorm:"column:uin" json:"uin"`
	Name        string `gorm:"column:name" json:"name"`
	Password    string `gorm:"column:password" json:"-"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
	Perm        int32  `gorm:"column:perm" json:"perm"`
}

func (obj *User) TableName() string {
	return "User"
}
