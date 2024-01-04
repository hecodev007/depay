package main

import (
	"depay/com"
	"depay/config"
	"depay/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt/rsa"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLog "gorm.io/gorm/logger"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func PostForm(uri, method, token string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, uri+ParseToStr(param), nil)
	req.Header.Add("signature", token)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	router.ServeHTTP(w, req)

	return w
}
func ParseToStr(mp map[string]interface{}) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + fmt.Sprintf("%v", val)
	}
	if len(values) == 0 {
		return ""
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

//
//priv: MIICXQIBAAKBgQDhhCkKiF+CoxwohxlM6R39AxbVWpwn7aqNOeRDXm21wZss/2OL/jWbHm1VHqvyjDOcYoO59B5MbEwZdx7H6BFS3l1aG1b9WE6C4t4FEgIZBsxa0u4RNsxa+nCt2YMsK3nDJjBMTuaFSLgueaWnhKK0xM9b+j069BBLcWXoCMsnJQIDAQABAoGALJs6jvXIhouCz1VhoL/fiaMpygvBJXiyKnsP9m9gHfpsmirt4svmiIctMw/9DN3Ee6NU0NxDffRR3RudwAbcHfwGVV68jXguHpAJsd00kz7orZlLHbPwDQnn83LVF7cgHDTwVZWc1mpzGh6UZpnwR3l04vhePOkYSoRT3y3jm6ECQQDwcAyU7fhAAcPHVq54RCzE2VgHWsIX0qr1TuZTacVwzTS6KLcqRcujI+022poKays9Ej8UpS5+Aq+9IJTPJcuJAkEA8BzehGrLDy8g+mK06+MUp0zWPj2t+NiHwPDGvPg7kwrPN0VJT1DXhUeyzSkmlz/nBmD6Rc3580E4AF+4Br0LvQJBALW/vYMGr9WSf++7Mn9u6XiT4tsMXBOuB9UPI0SCe+Fc/TKLfInT4K8dhT8l17Nwd2re1BhDFPXkCfwpGNPNeiECQGXf+dE49lrE5jsV8ik7OaIaCbRyuwOf60lDXy8CK1Sh+3U54nbSl/6mgwhk80itBjpAny9Ky0gYXcha1FuXjgkCQQDUPxa8zSPyj9ygaON5R/0YdqcLAhzu3hS+uJSRq2VIvTG+lFW42VZ+snFnTZqkM7rzPx9EVM0pCvWjCY9be1CI
//public: MIGJAoGBAOGEKQqIX4KjHCiHGUzpHf0DFtVanCftqo055ENebbXBmyz/Y4v+NZsebVUeq/KMM5xig7n0HkxsTBl3HsfoEVLeXVobVv1YToLi3gUSAhkGzFrS7hE2zFr6cK3ZgywrecMmMExO5oVIuC55paeEorTEz1v6PTr0EEtxZegIyyclAgMBAAE=

func TestGenOrder(t *testing.T) {

	globalrouter := Router()
	var w *httptest.ResponseRecorder
	param := make(map[string]interface{})
	param["merchant_address"] = "0x27Db22bfF368c4C42B2AF6f84De854757DefF7E7"
	param["merchant_id"] = 10001
	param["user_address"] = "0x96216849c49358B10257cb55b28eA603c874b05E"
	param["user_id"] = 10002
	param["usdt_amount"] = 8888
	token := "0x27Db22bfF368c4C42B2AF6f84De854757DefF7E7" + "&" + "0x96216849c49358B10257cb55b28eA603c874b05E"
	base64Sign, err := rsa.RsaSignBase64([]byte(token), "MIICXQIBAAKBgQDhhCkKiF+CoxwohxlM6R39AxbVWpwn7aqNOeRDXm21wZss/2OL/jWbHm1VHqvyjDOcYoO59B5MbEwZdx7H6BFS3l1aG1b9WE6C4t4FEgIZBsxa0u4RNsxa+nCt2YMsK3nDJjBMTuaFSLgueaWnhKK0xM9b+j069BBLcWXoCMsnJQIDAQABAoGALJs6jvXIhouCz1VhoL/fiaMpygvBJXiyKnsP9m9gHfpsmirt4svmiIctMw/9DN3Ee6NU0NxDffRR3RudwAbcHfwGVV68jXguHpAJsd00kz7orZlLHbPwDQnn83LVF7cgHDTwVZWc1mpzGh6UZpnwR3l04vhePOkYSoRT3y3jm6ECQQDwcAyU7fhAAcPHVq54RCzE2VgHWsIX0qr1TuZTacVwzTS6KLcqRcujI+022poKays9Ej8UpS5+Aq+9IJTPJcuJAkEA8BzehGrLDy8g+mK06+MUp0zWPj2t+NiHwPDGvPg7kwrPN0VJT1DXhUeyzSkmlz/nBmD6Rc3580E4AF+4Br0LvQJBALW/vYMGr9WSf++7Mn9u6XiT4tsMXBOuB9UPI0SCe+Fc/TKLfInT4K8dhT8l17Nwd2re1BhDFPXkCfwpGNPNeiECQGXf+dE49lrE5jsV8ik7OaIaCbRyuwOf60lDXy8CK1Sh+3U54nbSl/6mgwhk80itBjpAny9Ky0gYXcha1FuXjgkCQQDUPxa8zSPyj9ygaON5R/0YdqcLAhzu3hS+uJSRq2VIvTG+lFW42VZ+snFnTZqkM7rzPx9EVM0pCvWjCY9be1CI")
	if err != nil {
		fmt.Println(err)
		return
	}
	token = base64Sign
	//param["shareCode"] = "FN8AF7CP"
	urlLogin := "/genPayOrder"
	method := "POST"
	w = PostForm(urlLogin, method, token, param, globalrouter)

	body, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(body))
}

func TestGetPayOrder(t *testing.T) {
	globalrouter := Router()
	var w *httptest.ResponseRecorder
	//test post

	param := make(map[string]interface{})
	param["order_id"] = 1734452722045816832
	param["msg"] = "0x96216849c49358B10257cb55b28eA603c874b05E"
	token := com.Signature("0x96216849c49358B10257cb55b28eA603c874b05E")
	urlLogin := "/getPayOrder"
	method := "GET"
	w = PostForm(urlLogin, method, token, param, globalrouter)

	body, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(body))

}
func TestGetOrderStatus(t *testing.T) {
	globalrouter := Router()
	var w *httptest.ResponseRecorder
	//test post

	param := make(map[string]interface{})
	param["order_id"] = 1729706972187463680
	param["msg"] = "0x96216849c49358B10257cb55b28eA603c874b05E"
	token := com.Signature("0x96216849c49358B10257cb55b28eA603c874b05E")
	urlLogin := "/getOrderStatus"
	method := "GET"
	w = PostForm(urlLogin, method, token, param, globalrouter)

	body, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(body))

}

func TestCancelOrder(t *testing.T) {
	globalrouter := Router()
	var w *httptest.ResponseRecorder
	//test post

	param := make(map[string]interface{})
	param["order_id"] = 1729706972187463680
	param["msg"] = "0x96216849c49358B10257cb55b28eA603c874b05E"
	token := com.Signature("0x96216849c49358B10257cb55b28eA603c874b05E")
	urlLogin := "/cancelPayOrder"
	method := "POST"
	w = PostForm(urlLogin, method, token, param, globalrouter)

	body, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(body))

}

func TestDB(t *testing.T) {

	cfgPath := "./config/config.toml"

	if err := conf.Init(cfgPath); err != nil {
		panic(err)
	}
	fmt.Println(config.GlobalConf.Chain)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.HostPort,
		conf.Mysql.DBName)

	//全局使用，定义成共有的
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gLog.Default.LogMode(gLog.Info),
	})

	if err != nil {
		//log.Error(err)
		return
	} else {
		model.DB = db
	}

	order := &model.PayOrder{}
	if err := model.DB.Where("order_id=?", 1734452722045816832).First(order).Error; err != nil {

	}

	order.Chain = config.GlobalConf.Chain

	if err := model.DB.Where("id=?", order.Id).Save(order).Error; err != nil {

		fmt.Println("[PayOrderEvent] 修改支付订单错误：", order.OrderId)

	}

}
