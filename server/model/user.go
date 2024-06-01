package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`

	Uin         string `gorm:"column:uint;type:varchar(64);unique" json:"uin,omitempty"`
	Name        string `gorm:"column:name;type:varchar(20)" json:"name,omitempty"`
	Password    string `gorm:"column:password;type:varchar(20)" json:"-"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20);unique" json:"phone_number,omitempty"`
	Perm        int32  `gorm:"column:perm;type:int" json:"perm,omitempty"`
	Email       string `gorm:"column:email;type:varchar(20);unique" json:"email,omitempty"`
}

func (obj *User) TableName() string {
	return "User"
}
