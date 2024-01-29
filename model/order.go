package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// ico1688.info
type PayOrder struct {
	Id              int64           `json:"id"`
	OrderId         int64           `json:"order_id"`
	MerchantOrder   string          `json:"merchant_order"`
	UsdtAmount      decimal.Decimal `json:"usdt_amount"`
	PayedUsdt       decimal.Decimal `json:"payed_usdt"`
	SwapAmount      decimal.Decimal `json:"swap_amount"`
	TokenAmount     decimal.Decimal `json:"token_amount"`
	TokenAddress    string          `json:"token_address"`
	Chain           string          `json:"chain"`
	Status          int             `json:"status"`
	Notifyed        int             `json:"notifyed"`
	MerchantAddress string          `json:"merchant_address"`
	MerchantId      int64           `json:"merchant_id"`
	UserAddress     string          `json:"user_address"`
	UserId          int64           `json:"user_id"`
	SuccessUrl      string          `json:"success_url"`
	CancelUrl       string          `json:"cancel_url"`
	TxId            string          `json:"tx_id"`
	RecAddress      string          `json:"rec_address"`
	CreateTime      time.Time       `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime      time.Time       `json:"update_time" gorm:"comment:更新时间"`
}

func (PayOrder) TableName() string {
	return "pay_order"
}

const (
	UNPAY = iota
	PAYING
	PAYED
	TIMEOUT
	CANCEL
	PART
)
