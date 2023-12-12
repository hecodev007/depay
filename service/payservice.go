package service

import (
	"depay/com"
	"depay/model"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/wumansgy/goEncrypt/rsa"
	"net/http"
	"time"
)

type Service struct {
	Node *snowflake.Node
}
type PayOrderReq struct {
	MerchantAddress string          `form:"merchant_address" json:"merchant_address" binding:"required"`
	MerchantId      int64           `form:"merchant_id"  json:"merchant_id" binding:"required"`
	UserAddress     string          `form:"user_address" json:"user_address" binding:"required"`
	UserId          int64           `form:"user_id" json:"user_id"`
	UsdtAmount      decimal.Decimal `form:"usdt_amount" json:"usdt_amount" binding:"required"`
}

func (s *Service) GenPayOrder(c *gin.Context) {
	req := PayOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	signature := c.GetHeader("signature")
	msg := req.MerchantAddress + "&" + req.UserAddress
	PublicKey := "MIGJAoGBAOGEKQqIX4KjHCiHGUzpHf0DFtVanCftqo055ENebbXBmyz/Y4v+NZsebVUeq/KMM5xig7n0HkxsTBl3HsfoEVLeXVobVv1YToLi3gUSAhkGzFrS7hE2zFr6cK3ZgywrecMmMExO5oVIuC55paeEorTEz1v6PTr0EEtxZegIyyclAgMBAAE="

	res := rsa.RsaVerifySignBase64([]byte(msg), signature, PublicKey)
	if !res {
		c.JSON(http.StatusOK, gin.H{"msg": "signature is error", "code": 1})
		return
	}
	id := s.Node.Generate().Int64()
	strId := fmt.Sprintf("%v", id)

	strId = strId[len(strId)-15:]
	dec, _ := decimal.NewFromString(strId)
	order := &model.PayOrder{
		OrderId:         dec.IntPart(),
		UserId:          req.UserId,
		UserAddress:     req.UserAddress,
		MerchantAddress: req.MerchantAddress,
		MerchantId:      req.MerchantId,
		Status:          model.UNPAY,
		UsdtAmount:      req.UsdtAmount,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}
	if err := model.DB.Create(order).Error; err != nil {
		log.Println("gen pay order err:", err)
		c.JSON(http.StatusOK, gin.H{"msg": "db err", "code": 1})
		return
	}
	fmt.Println("gen order rsp:", order.OrderId)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "code": 0, "order_id": order.OrderId})
	return
}
func (s *Service) TestApi(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

// getPayOrder?order_id=100012&msg=address
type GetOrderReq struct {
	OrderId int64  `form:"order_id" json:"order_id" binding:"required"`
	Msg     string `form:"msg" json:"msg" binding:"required"`
}

func (s *Service) GetPayOrder(c *gin.Context) {
	req := GetOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	order := &model.PayOrder{}
	if err := model.DB.Where("order_id=?", req.OrderId).First(order).Error; err != nil {
		log.Println("query db error:", req.OrderId, err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db error"})
		return
	}
	if common.HexToAddress(order.UserAddress) != common.HexToAddress(req.Msg) {
		log.Println("req address is not user address:", req.Msg, order.UserAddress)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param error"})
		return
	}
	signature := c.GetHeader("signature")
	b := com.ValidSigner([]byte(signature), req.Msg, common.HexToAddress(req.Msg))
	if !b {
		log.Println("valid signature:", req.OrderId, req.Msg)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "valid signature！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "merchant_address": order.MerchantAddress, "merchant_id": order.MerchantId, "usdt_amount": order.UsdtAmount, "order_id": order.OrderId})
}

type GetOrderStatusReq struct {
	OrderId int64  `form:"order_id"  json:"order_id" binding:"required"`
	Msg     string `form:"msg"  json:"msg" binding:"required"`
}

func (s *Service) GetOrderStatus(c *gin.Context) {
	req := GetOrderStatusReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	order := &model.PayOrder{}
	if err := model.DB.Where("order_id=?", req.OrderId).First(order).Error; err != nil {
		log.Println("query db error:", req.OrderId)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db error"})
		return
	}
	if common.HexToAddress(order.UserAddress) != common.HexToAddress(req.Msg) {
		log.Println("req address is not user address:", req.Msg, order.UserAddress)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param error"})
		return
	}
	signature := c.GetHeader("signature")
	b := com.ValidSigner([]byte(signature), req.Msg, common.HexToAddress(req.Msg))
	if !b {
		log.Println("valid signature:", req.OrderId, req.Msg)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "valid signature！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "status": order.Status, "order_id": order.OrderId})

}

type CancelOrderReq struct {
	OrderId int64  `form:"order_id"  json:"order_id" binding:"required"`
	Msg     string `form:"msg"   json:"msg" binding:"required"`
}

func (s *Service) CancelPayOrder(c *gin.Context) {
	req := CancelOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param err！"})
		return
	}
	order := &model.PayOrder{}
	if err := model.DB.Where("order_id=?", req.OrderId).First(order).Error; err != nil {
		log.Println("query db error:", req.OrderId)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "db error"})
		return
	}
	if common.HexToAddress(order.UserAddress) != common.HexToAddress(req.Msg) {
		log.Println("req address is not user address:", req.Msg, order.UserAddress)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "param error"})
		return
	}
	signature := c.GetHeader("signature")
	b := com.ValidSigner([]byte(signature), req.Msg, common.HexToAddress(req.Msg))
	if !b {
		log.Println("valid signature:", req.OrderId, req.Msg)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "valid signature！"})
		return
	}
	order.Status = model.CANCEL
	if err := model.DB.Where("id=?", order.Id).Save(order).Error; err != nil {
		log.Println("update status error:", order.OrderId)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "update status error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})

}
