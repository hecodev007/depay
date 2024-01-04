package main

import (
	"depay/middleware"
	"depay/middleware/auth"
	"depay/service"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLog "gorm.io/gorm/logger"
	"os"
	"os/signal"

	"depay/config"
	"depay/loger"
	"depay/model"
	"time"
)

func init() {
	model.GoCache = cache.New(10*time.Hour, 60*time.Second)
	//	com.Mu = sync.RWMutex{}
	//loger.InitLog(".", "reward", "")
}

var conf = &config.Config{}

// redis-cli -h abeatsdev.fss19b.ng.0001.use2.cache.amazonaws.com -p 6379
func main() {
	//pri := flag.String("p", "", "owner private")
	env := flag.Int64("e", 0, "owner private")
	flag.Parse()

	loger.InitLog(".", "admin", "")
	cfgPath := "./config/config.toml"
	if *env == 1 {
		cfgPath = "./config/config_pro.toml"
	}

	//loger.InitLog(".", "reward", "")
	//cfgPath := "../config/config.toml"
	if err := conf.Init(cfgPath); err != nil {
		panic(err)
	}

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
		log.Error(err)
		return
	} else {
		model.DB = db
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := service.Service{
		Node: node,
	}
	r := gin.Default()
	auth.NoCheckUrl = append(auth.NoCheckUrl, "/admin/register")
	auth.NoCheckUrl = append(auth.NoCheckUrl, "/admin/login")
	auth.NoCheckUrl = append(auth.NoCheckUrl, "/admin/getEmailCode")
	auth.NoCheckUrl = append(auth.NoCheckUrl, "/")
	r.Use(middleware.Cors())
	//	auth.NoCheckUrl = append(auth.NoCheckUrl, "/register")
	r.Use(auth.JWTAuth())
	router := r.Group("admin")
	router.POST("/login", s.Login)
	router.POST("/register", s.RegUser)
	//router.POST("/setWalletAddress", s.SetWalletAddress)
	router.POST("/setWebHook", s.SetWebHook)
	router.POST("/changePwd", s.ChangePwd)
	router.POST("/setCoin", s.SetCoin)
	router.POST("/logout", s.LogOut)
	router.POST("/getEmailCode", s.GetEmailCode)
	router.GET("/getMerchantInfo", s.GetMerchantInfo)
	router.GET("/getCoinInfo", s.GetCoinInfo)
	router.GET("/getPayOrder", s.GetPayOrders)
	router.GET("/getRequestLog", s.GetRequestLog)
	router.POST("/delAddress", s.DelCoin)

	fmt.Println("start serve。。。")

	r.Run(conf.AdminPort)
	// 捕捉退出信号
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

}

// timeStamp
func Router() *gin.Engine {
	loger.InitLog(".", "reward", "")

	model.GoCache = cache.New(10*time.Hour, 60*time.Second)

	conf := config.New()
	cfgPath := "./config/config.toml"
	if err := conf.Init(cfgPath); err != nil {
		panic(err)
	}
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
		log.Error(err)
	} else {
		model.DB = db
	}
	r := gin.Default()
	//s := &service.Service{
	//}
	//<-done

	//r.GET("/getWebsiteInfo", getWebsiteInfo)

	return r
}
