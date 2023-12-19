package model

import (
	"time"
)

type MerchantAddress struct {
	Id         string    `json:"id"`
	Address    string    `json:"address"`
	MerchantId int64     `json:"merchant_id"`
	CreateTime time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime time.Time `json:"update_time"`
	Coin       string    `json:"coin"`
	Chain      string    `json:"chain"`
}

func (MerchantAddress) TableName() string {
	return "merchant_address"
}
