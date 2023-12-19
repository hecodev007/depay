package model

import (
	"time"
)

type Merchant struct {
	Id           string    `json:"id"`
	MerchantId   int64     `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	CreateTime   time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime   time.Time `json:"update_time"`
	SecretKey    string    `json:"secret_key"`
	PublicKey    string    `json:"public_key"`
	WebHook      string    `json:"web_hook"`
}

func (Merchant) TableName() string {
	return "merchant"
}
