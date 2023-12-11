package service

import (
	"depay/middleware/auth"
	"depay/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AddMerchantReq struct {
	MerchantAddress string `json:"merchant_address"`
	MerchantId      int64  `json:"merchant_id"`
	MerchantName    string `json:"merchant_name"`
}

func (s *Service) AddMerchant(c *gin.Context) {
	req := AddMerchantReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	merchant := &model.Merchant{
		MerchantId:      s.Node.Generate().Int64(),
		MerchantAddress: req.MerchantAddress,
		MerchantName:    req.MerchantName,
	}
	if err := model.DB.Create(merchant).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "insert db  err！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

type AddAdminReq struct {
	Email string `json:"email" binding:"required"`
	Pwd   string `json:"pwd" binding:"required"`
	Level int    `json:"level"`
}

func (s *Service) RegUser(c *gin.Context) {
	req := AddAdminReq{}
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
	if claims.Level != 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "no access to add user err！"})
		return
	}
	admin := &model.Admin{
		UserName: req.Email,
		Pwd:      req.Pwd,
		Level:    req.Level,
	}
	if err := model.DB.Create(admin).Error; err != nil {
		log.Println("add user err:", err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": " add user err！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": " success"})
	return
}

type AdminLogReq struct {
	UserName int64  `json:"user_name" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
}

func (s *Service) Login(c *gin.Context) {
	req := AdminLogReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	admin := &model.Admin{}
	if err := model.DB.Where("user_name").First(admin).Error; err != nil {
		log.Error("admin select user err:", req.UserName)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	if admin.Pwd != req.Pwd {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "pwd err！"})
		return
	}
	auth.GenerateToken(c, *admin)
}
