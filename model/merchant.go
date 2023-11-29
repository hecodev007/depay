package model

import (
	"time"
)

type Merchant struct {
	Id              string    `json:"id"`
	MerchantAddress string    `json:"merchant_address"`
	MerchantId      int64     `json:"merchant_id"`
	MerchantName    string    `json:"merchant_name"`
	CreateTime      time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime      time.Time `json:"update_time"`
}

func (Merchant) TableName() string {
	return "merchant"
}
