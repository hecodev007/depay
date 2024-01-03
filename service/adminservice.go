package service

import (
	"depay/com"
	"depay/middleware/auth"
	"depay/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/wumansgy/goEncrypt/rsa"
	"net/http"
	"strings"
	"time"
)

func (s *Service) DelCoin(c *gin.Context) {
	req := DelCoinReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		log.Error("parse token ", err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！" + err.Error()})
		return
	}

	if err := model.DB.Model(&model.MerchantAddress{}).Where("merchant_id=? and chain=? and address = ?", claims.MerchantId, req.Chain, req.Address).Delete(&model.MerchantAddress{}).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "insert db  err！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

type GetEmailCodeReq struct {
	Email string `json:"email" form:"email" binding:"required"`
}

func (s *Service) GetEmailCode(c *gin.Context) {
	req := GetEmailCodeReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	err := com.SendMail(req.Email)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "send email err！"})
		return
	}
	code, _ := model.GoCache.Get(req.Email)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "email_code": code})
}

func (s *Service) LogOut(c *gin.Context) {

	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		log.Error("parse token ", err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！" + err.Error()})
		return
	}
	model.GoCache.Delete(claims.UserName)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

type DelCoinReq struct {
	//Coin    string `json:"coin" form:"coin" binding:"required"`
	Chain   string `json:"chain" form:"chain" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}
type SetCoinReq struct {
	Coin    string `json:"coin" form:"coin" binding:"required"`
	Chain   string `json:"chain" form:"chain" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}

func (s *Service) SetCoin(c *gin.Context) {
	req := SetCoinReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		log.Error("parse token ", err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！" + err.Error()})
		return
	}

	if err := model.DB.Model(&model.MerchantAddress{}).Where("merchant_id=? and chain=? and address = ?", claims.MerchantId, req.Chain, req.Address).Delete(&model.MerchantAddress{}).Error; err != nil {
		log.Info(err)
	}
	coins := strings.Split(req.Coin, ",")
	for _, v := range coins {
		merchant := &model.MerchantAddress{}
		merchant.Address = req.Address
		merchant.MerchantId = claims.MerchantId
		merchant.Coin = v
		merchant.Chain = req.Chain
		merchant.CreateTime = time.Now()
		merchant.UpdateTime = time.Now()
		if err := model.DB.Create(merchant).Error; err != nil {
			log.Error(err)
		}
		//c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
		//return

	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

type ChangePwdReq struct {
	Email string `json:"email" form:"email"  binding:"required"`
	Pwd   string `json:"pwd"  form:"pwd" binding:"required"`
	Code  string `json:"code"  form:"code" binding:"required"`
}

func (s *Service) ChangePwd(c *gin.Context) {
	req := ChangePwdReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	//token := c.GetHeader("token")
	//j := auth.NewJWT()
	//// parseToken 解析token包含的信息
	//claims, err := j.ParseToken(token)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
	//	return
	//}
	code, is := model.GoCache.Get(req.Email)
	cd := code.(string)
	if !is || req.Code != cd {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "code is err！"})
		return
	}
	merchant := &model.Admin{}
	if err := model.DB.Model(merchant).Where("user_name=?", req.Email).Update("pwd", req.Pwd).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token  err！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

type SetWebHookReq struct {
	Url string `json:"url"  form:"url" binding:"required"`
}

func (s *Service) SetWebHook(c *gin.Context) {
	req := SetWebHookReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
		return
	}
	merchant := &model.Merchant{}
	if err := model.DB.Model(merchant).Where("merchant_id=?", claims.MerchantId).First(merchant).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token  err！"})
		return
	}

	merchant.WebHook = req.Url
	if err := model.DB.Save(merchant).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "insert db  err！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

type GetRequestLogReq struct {
	PageSize  int    `json:"page_size"  form:"page_size" binding:"required"`
	PageIndex int    `json:"page_index"  form:"page_index" binding:"required"`
	StartTime string `json:"start_time"  form:"start_time" binding:"required"`
	EndTime   string `json:"end_time"  form:"end_time" binding:"required"`
	Direct    int    `json:"direct" form:"direct" `
}

func (s *Service) GetRequestLog(c *gin.Context) {
	req := GetRequestLogReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
		return
	}

	orders := make([]model.RequestLog, 0)
	total := int64(0)
	if err := model.DB.Order("id desc").Where("merchant_id=? and direct=? and create_time >= ? and create_time <= ?", claims.MerchantId, req.Direct, req.StartTime, req.EndTime).Model(&model.RequestLog{}).Count(&total).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusOK, gin.H{"code": 2, "msg": "no data"})
			return
		}

		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db err！"})
		return
	}

	if err := model.DB.Order("id desc").Where("merchant_id=? and  direct=? and create_time >= ? and create_time <= ?", claims.MerchantId, req.Direct, req.StartTime, req.EndTime).Model(&model.RequestLog{}).Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize).Find(&orders).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db err！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "total": total, "list": orders})
}

type PayOrdersReq struct {
	PageSize  int    `json:"page_size"  form:"page_size" binding:"required"`
	PageIndex int    `json:"page_index"  form:"page_index" binding:"required"`
	StartTime string `json:"start_time"  form:"start_time" binding:"required"`
	EndTime   string `json:"end_time"  form:"end_time" binding:"required"`
	Status    int    `json:"status"  form:"status" binding:"required"`
}

func (s *Service) GetPayOrders(c *gin.Context) {
	req := PayOrdersReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
		return
	}

	orders := make([]model.PayOrder, 0)
	total := int64(0)
	if req.Status == -1 {

		if err := model.DB.Where("merchant_id=? and create_time >= ? and create_time <= ?", claims.MerchantId, req.StartTime, req.EndTime).Model(&model.PayOrder{}).Count(&total).Error; err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token  err！"})
			return
		}

		if err := model.DB.Order("id desc").Where("merchant_id=? and create_time >= ? and create_time <= ?", claims.MerchantId, req.StartTime, req.EndTime).Model(&model.PayOrder{}).Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize).Find(&orders).Error; err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token  err！"})
			return
		}
	} else if req.Status == 2 {
		if err := model.DB.Where("merchant_id=? and status=2 and create_time >= ? and create_time <= ?", claims.MerchantId, req.StartTime, req.EndTime).Model(&model.PayOrder{}).Count(&total).Error; err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db  err！"})
			return
		}

		if err := model.DB.Order("id desc").Where("merchant_id=? and  status=2 and create_time >= ? and create_time <= ?", claims.MerchantId, req.StartTime, req.EndTime).Model(&model.PayOrder{}).Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize).Find(&orders).Error; err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db  err！"})
			return
		}
	} else {
		if err := model.DB.Where("merchant_id=? and status !=2 and create_time >= ? and create_time <= ?", claims.MerchantId, req.StartTime, req.EndTime).Model(&model.PayOrder{}).Count(&total).Error; err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db  err！"})
			return
		}

		if err := model.DB.Order("id desc").Where("merchant_id=? and  status !=2 and create_time >= ? and create_time <= ?", claims.MerchantId, req.StartTime, req.EndTime).Model(&model.PayOrder{}).Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize).Find(&orders).Error; err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db  err！"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "total": total, "list": orders})
}

func (s *Service) GetMerchantInfo(c *gin.Context) {

	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
		return
	}
	merchant := &model.Merchant{}
	if err := model.DB.Model(merchant).Where("merchant_id=?", claims.MerchantId).First(merchant).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token  err！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "merchant_info": merchant})
}

//	type CoinInfoReq struct {
//		Chain string `json:"chain"  form:"chain" binding:"required"`
//	}
type CoinChain struct {
	Chain   string
	Address string
	Coin    []string
}

func (s *Service) GetCoinInfo(c *gin.Context) {
	//req := CoinInfoReq{}
	//if err := c.ShouldBind(&req); err != nil {
	//	log.Error(err)
	//	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
	//	return
	//}
	token := c.GetHeader("token")
	j := auth.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
		return
	}
	merchants := make([]model.MerchantAddress, 0)
	if err := model.DB.Model(&model.MerchantAddress{}).Where("merchant_id=? ", claims.MerchantId).Find(&merchants).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token  err！"})
		return
	}
	list := make(map[string]CoinChain)
	for _, v := range merchants {
		if a, ok := list[v.Address+","+v.Chain]; ok {
			a.Coin = append(a.Coin, v.Coin)
			list[v.Address+","+v.Chain] = a
		} else {
			coin := make([]string, 0)
			coin = append(coin, v.Coin)
			list[v.Address+","+v.Chain] = CoinChain{
				v.Chain, v.Address, coin,
			}
		}
	}
	fmt.Println("list:", list)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "list": list})
}

type AddAdminReq struct {
	Email string `json:"email" form:"email"  binding:"required"`
	Pwd   string `json:"pwd"  form:"pwd" binding:"required"`
	Code  string `json:"code" form:"code"  binding:"required"`
}

func (s *Service) RegUser(c *gin.Context) {
	req := AddAdminReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	//token := c.GetHeader("token")
	//j := auth.NewJWT()
	//// parseToken 解析token包含的信息
	//claims, err := j.ParseToken(token)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "token err！"})
	//	return
	//}
	//if claims.Level != 1 {
	//	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "no access to add user err！"})
	//	return
	//}
	code, b := model.GoCache.Get(req.Email)
	a, _ := code.(string)
	if !b || req.Code != a {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "code error"})
		return
	}
	id := s.Node.Generate().Int64()
	strId := fmt.Sprintf("%v", id)
	strId = strId[len(strId)-11:]
	dec, _ := decimal.NewFromString(strId)
	admin := &model.Admin{
		UserName:   req.Email,
		Pwd:        req.Pwd,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Enable:     1,
		MerchantId: dec.IntPart(),
	}
	tx := model.DB.Begin()
	if err := tx.Create(admin).Error; err != nil {
		log.Println("add user err:", err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": " add user err！"})
		tx.Rollback()
		return
	}
	rsaBase64Key, err := rsa.GenerateRsaKeyBase64(1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	merchant := &model.Merchant{
		MerchantId: admin.MerchantId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		SecretKey:  rsaBase64Key.PrivateKey,
		PublicKey:  rsaBase64Key.PublicKey,
	}
	if err := tx.Create(merchant).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "insert db  err！"})
		tx.Rollback()
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": " success"})
	return
}

type AdminLogReq struct {
	Email string `json:"email" form:"email" binding:"required"`
	Pwd   string `json:"pwd" form:"pwd" binding:"required"`
}

func (s *Service) Login(c *gin.Context) {
	req := AdminLogReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	admin := &model.Admin{}
	if err := model.DB.Where("user_name=?", req.Email).First(admin).Error; err != nil {
		log.Error("admin select user err:", req.Email)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	if admin.Pwd != req.Pwd {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "pwd err！"})
		return
	}
	auth.GenerateToken(c, *admin)
}
