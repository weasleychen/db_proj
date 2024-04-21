package model

import "gorm.io/gorm"

type Dish struct {
	gorm.Model

	Name     string  `gorm:"column:name"`
	Price    float64 `gorm:"column:price"`
	Discount float64 `gorm:"column:discount"`
	Detail   string  `gorm:"column:detail"`
}

func (item *Dish) TableName() string {
	return "Dish"
}
