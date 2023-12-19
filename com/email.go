package com

import (
	"bytes"
	"depay/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	//"time"
)

//func SendMail(email string) error {
//	m := gomail.NewMessage()
//	m.SetAddressHeader("From", "markjeff1688@gmail.com", "mark jeff") // 发件人邮箱，发件人名称
//	m.SetHeader("To",                                      // 收件人
//		m.FormatAddress(email, "Receiver"),
//	)
//	//validCode := GenValidateCode(6)
//	msg:="789534"
//	m.SetHeader("Subject", "验证邮件")                                                                                                                   // 主题
//	m.SetBody("text/html", fmt.Sprintf("%s%s%s", "尊敬的用户您好！<br>验证码：", msg, "验证码有效时间：10分钟。请勿向任何人包括客服提供验证码。<br>感谢您选择QueerPay<br><br>【QueerPay】")) // 正文
//
//	d := gomail.NewDialer("smtp.gmail.com", 465, "markjeff1688@gmail.com", "ico1688.info") // 发送邮件服务器、端口、发件人账号、发件人密码
//	if err := d.DialAndSend(m); err != nil {
//		fmt.Println(err)
//		return errors.New("发送邮件出错！")
//	}
//	//store.GoCache.Set(email, validCode, 10*time.Minute)
//	return nil
//}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type To struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Email struct {
	Sender      Sender `json:"sender"`
	To          []To   `json:"to"`
	Subject     string `json:"subject"`
	HtmlContent string `json:"htmlContent"`
}

func SendData(data, _to string) error {
	//host := fmt.Sprintf(url2, page, pageSize)
	//	fmt.Println(host)
	//body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	body := Email{
		Sender: Sender{
			Name:  "rns pay",
			Email: "notification@email.abeats.com",
		},
		To: []To{{
			Name:  "W",
			Email: _to,
		}},
		Subject: "Validate Code",
	}
	html := fmt.Sprintf("<html><head></head><body><p>your code is:</p>%v.</p></body></html>", data)
	body.HtmlContent = html
	dt, _ := json.Marshal(body)
	fmt.Println("send:", string(dt))
	request, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(dt))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-key", "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")
	request.Header.Set("accept", "application/json")
	//异常捕捉
	if err != nil {
		fmt.Println(err)
		return err
	}
	client := &http.Client{}
	//处理返回结果
	response, _ := client.Do(request)
	//关闭流
	defer response.Body.Close()
	//检出结果集
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll failed ,err:%v", err)
		return err

	}

	//list:=make([]Data,0)
	//fmt.Println("Data:", string(rest))
	//r := DT{}
	//json.Unmarshal(body, &r)
	//fmt.Println("rsp:", r)
	//return r, err
	return nil
}
func SendMail(_to string) error {
	validCode := GenValidateCode(6)
	fmt.Println("验证码：", validCode)

	err := SendData(validCode, _to)
	if err != nil {
		return err
	}
	model.GoCache.Set(_to, validCode, 10*time.Minute)
	return nil
}

//func SendMail(_to string) error {
//	from := "wonderAbeats@gmail.com"
//	password := "lftjbwzrwvhyhdun"
//	to := []string{
//		_to,
//	}
//	validCode :=GenValidateCode(6)
//	message := []byte(validCode)
//	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
//
//	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
//	if err != nil {
//		fmt.Println("SendMail:",err)
//		return err
//	}
//	fmt.Println("Email Sent!")
//	model.GoCache.Set(_to, validCode, 10*time.Minute)
//	return nil
//}
