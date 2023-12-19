package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

var NoCheckUrl []string

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		b := noCheck(c)
		if b {
			c.Next()
		} else {
			token := c.Request.Header.Get("token")
			if token == "" {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "请求未携带token，无权限访问",
				})
				c.Abort()
				return
			}

			log.Print("get token: ", token)

			j := NewJWT()
			// parseToken 解析token包含的信息
			_, err := j.ParseToken(token)
			if err != nil {
				if err == TokenExpired {
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "授权已过期",
					})
					c.Abort()
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  err.Error(),
				})
				c.Abort()
				return
			}
			//db.GoCache.Set("userInfo",data,60*time.Minute)
			//db.GoCache.Set("userId",data.UserId,60*time.Minute)
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			//c.Set("claims", claims)
		}

	}
}

func noCheck(c *gin.Context) bool {
	path := c.Request.URL.Path
	for _, p := range NoCheckUrl {
		if path == p || strings.Contains(path, "swagger") {
			return true
		}
	}
	return false
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserName   string `json:"user_name"`
	UserId     int64  `json:"user_id"`
	MerchantId int64  `json:"merchant_id"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
