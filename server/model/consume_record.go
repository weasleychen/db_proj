package model

import "gorm.io/gorm"

type ConsumeRecordJson struct {
	TableId       int32   `json:"table_id"`
	Uin           string  `json:"uin"`
	OrderedDishes []Dish  `json:"orderedDishes"`
	Discount      float64 `json:"discount"`
	OriginPrice   float64 `json:"origin_price"`
	FinalPrice    float64 `json:"final_price"`
	CompleteTableTime int64 `json:"complete_table_time"`
}

type ConsumeRecord struct {
	gorm.Model
	Data string `gorm:"type:text"`
}

func (it *ConsumeRecord) TableName() string {
	return "ConsumeRecord"
}
