package model

import "gorm.io/gorm"

type Dish struct {
	gorm.Model `json:"gorm.Model"`

	Name     string  `gorm:"column:name; type:varchar(10)" json:"name,omitempty"`
	Price    float64 `gorm:"column:price; type:double" json:"price,omitempty"`
	Discount float64 `gorm:"column:discount; type:double" json:"discount,omitempty"`
	Detail   string  `gorm:"column:detail; type:text" json:"detail,omitempty"`
}

func (item *Dish) TableName() string {
	return "Dish"
}
