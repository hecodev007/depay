package model

import "time"

type Admin struct {
	Id         int64     `json:"id"`
	UserName   string    `json:"user_name"`
	Pwd        string    `json:"pwd"`
	Enable     int       `json:"enable"`
	MerchantId int64     `json:"merchant_id"`
	CreateTime time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime time.Time `json:"update_time"`
}

func (Admin) TableName() string {
	return "admin"
}
