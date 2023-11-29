package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type SubOrder struct {
	Id              string          `json:"id"`
	OrderId         int64           `json:"orderId"`
	Product         int64           `json:"product"`
	Amount          decimal.Decimal `json:"amount"`
	TokenAddress    string          `json:"token_address"`
	Status          int             `json:"status"`
	MerchantAddress string          `json:"merchant_address"`
	UserAddress     string          `json:"user_address"`
	MerchantId      int64           `json:"merchant_id"`
	CreateTime      time.Time       `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime      time.Time       `json:"update_time"`
}

func (SubOrder) TableName() string {
	return "sub_order"
}
