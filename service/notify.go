package service

import (
	"bytes"
	"depay/model"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Notice struct {
	MerchantOrder string `json:"merchant_order"`
	UsdtAmount    string `json:"usdt_amount"`
	PayedAmount   string `json:"payed_amount"`
	Status        int    `json:"status"` //2是成功
}

func Notify() {
	orders := make([]model.PayOrder, 0)
	if err := model.DB.Where("notifyed=? and status=?", 0, model.PAYED).Find(&orders).Error; err != nil {
		log.Println("query db error:", err)

		return
	}
	for _, v := range orders {
		merchant := &model.Merchant{}
		if err := model.DB.Model(merchant).Where("merchant_id=?", v.MerchantId).First(merchant).Error; err != nil {
			log.Error(err)
			return
		}
		notice := &Notice{
			MerchantOrder: v.MerchantOrder,
			Status:        v.Status,
			UsdtAmount:    v.UsdtAmount.String(),
			PayedAmount:   v.PayedUsdt.String(),
		}
		dt, _ := json.Marshal(notice)
		fmt.Println("notify:", string(dt))
		request, err := http.NewRequest("POST", merchant.WebHook, bytes.NewBuffer(dt))
		request.Header.Set("Content-Type", "application/json")
		//	request.Header.Set("api-key", "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")
		request.Header.Set("accept", "application/json")
		//request.Header.Set("Some-Custom-Name", "unique-id-1234")

		//异常捕捉
		if err != nil {
			fmt.Println(err)
			return
		}
		client := &http.Client{}
		//处理返回结果
		response, _ := client.Do(request)
		//关闭流
		defer response.Body.Close()
		//检出结果集
		resp, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("ioutil.ReadAll failed ,err:%v", err)
			return

		}

		logs := &model.RequestLog{
			Interface:  merchant.WebHook,
			Method:     "POST",
			Req:        string(dt),
			Rsp:        string(resp),
			Status:     response.StatusCode,
			MerchantId: v.MerchantId,
			CreateTime: time.Now(),
		}

		tx := model.DB.Begin()
		tx.Create(logs)
		if response.StatusCode == 200 {
			v.Notifyed = 1
			tx.Save(&v)
		}
		tx.Commit()

	}

}
