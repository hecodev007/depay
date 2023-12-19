package model

import "time"

type RequestLog struct {
	Id         int64     `json:"id" gorm:"type:int;comment:id"`
	Interface  string    `json:"interface"`
	Method     string    `json:"method"`
	Req        string    `json:"req"`
	Rsp        string    `json:"rsp"`
	CreateTime time.Time `json:"create_time"`
	MerchantId int64     `json:"merchantId"`
}

func (RequestLog) TableName() string {
	return "request_log"
}
