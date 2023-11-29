package model

import (
	"time"
)

type Subscribes struct {
	Id              string    `json:"id"`
	MerchantAddress string    `json:"merchant_address"`
	UserAddress     string    `json:"user_address"`
	Product         int64     `json:"product"`
	Period          int64     `json:"period"`
	CreateTime      time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime      time.Time `json:"update_time"`
}

func (Subscribes) TableName() string {
	return "subscribes"
}
