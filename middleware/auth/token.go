package auth

import (
	"depay/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/patrickmn/go-cache"

	"net/http"
	"time"
)

func GenerateToken(c *gin.Context, user model.Admin) {
	j := &JWT{
		[]byte("newtrekWang"),
	}
	claims := CustomClaims{
		user.UserName,
		user.Id,
		user.MerchantId,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),  // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 36000), // 过期时间 一小时
			Issuer:    "otcFrontUser",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	data := LoginResult{
		Token:      token,
		UserId:     user.Id,
		UserName:   user.UserName,
		MerchantId: user.MerchantId,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功！",
		"data": data,
	})
	model.GoCache.Set(user.UserName, 1, cache.NoExpiration)
	return
}

type LoginResult struct {
	Token      string `json:"token"`
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	MerchantId int64  `json:"merchant_id"`
}
