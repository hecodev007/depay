package model

import (
	"time"
)

type ProductInfo struct {
	Id              string `json:"id"`
	MerchantAddress string `json:"merchant_address"`
	Product         int64  `json:"product"`
	Period          int
	CreateTime      time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime      time.Time `json:"update_time"`
}

func (ProductInfo) TableName() string {
	return "product_info"
}

const (
	MONTH = iota + 1
	QUARTER
	HALFYEAR
	YEAR
)
