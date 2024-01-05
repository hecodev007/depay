package middleware

import (
	"bytes"
	"depay/model"
	"depay/service"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		//开始时间
		//	startTime := time.Now().UnixMicro()

		//处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		//结束时间
		//	endTime := time.Now()

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}
		logs := &model.RequestLog{
			Interface:  c.Request.RequestURI,
			Method:     c.Request.Method,
			Req:        c.Request.PostForm.Encode(),
			Rsp:        responseBody,
			Status:     c.Writer.Status(),
			Direct:     1,
			CreateTime: time.Now(),
		}

		if strings.Contains(logs.Interface, "genPayOrder") {

			req := service.PayOrderReq{}
			if err := c.ShouldBind(&req); err != nil {
				log.Println(err)
			}
			logs.MerchantId = req.MerchantId
			if err := model.DB.Create(logs); err != nil {
				log.Println(err)
			}
			model.DB.Create(logs)
		}

		////日志格式
		//accessLogMap := make(map[string]interface{})
		//
		//accessLogMap["request_time"] = startTime
		//accessLogMap["request_method"] = c.Request.Method
		//accessLogMap["request_uri"] = c.Request.RequestURI
		//accessLogMap["request_proto"] = c.Request.Proto
		//accessLogMap["request_ua"] = c.Request.UserAgent()
		//accessLogMap["request_referer"] = c.Request.Referer()
		//accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		//accessLogMap["request_client_ip"] = c.ClientIP()
		//
		//accessLogMap["response_time"] = endTime
		//accessLogMap["response_code"] = responseCode
		//accessLogMap["response_msg"] = responseMsg
		//accessLogMap["response_data"] = responseData
		//
		//accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)
		//
		//accessLogJson, _ := util.JsonEncode(accessLogMap)
		//
		//if f, err := os.OpenFile(config.AppAccessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		//	log.Println(err)
		//} else {
		//	f.WriteString(accessLogJson + "\n")
		//}
	}
}
