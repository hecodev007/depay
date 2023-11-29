package main

import (
	"depay/middleware/auth"
	"depay/service"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"github.com/vrecan/death"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLog "gorm.io/gorm/logger"

	"depay/config"
	"depay/loger"
	"depay/model"
	"syscall"
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
	flag.Parse()

	loger.InitLog(".", "reward", "")
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
	auth.NoCheckUrl = append(auth.NoCheckUrl, "/login")
	auth.NoCheckUrl = append(auth.NoCheckUrl, "/")
	//	auth.NoCheckUrl = append(auth.NoCheckUrl, "/register")
	r.Use(auth.JWTAuth())
	r.POST("/login", s.Login)
	r.GET("/register", s.RegUser)
	r.POST("/addMerchant", s.AddMerchant)

	fmt.Println("start serve。。。")

	r.Run(conf.Port)
	// 捕捉退出信号
	d := death.NewDeath(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL,
		syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGALRM)
	d.WaitForDeathWithFunc(func() {
		fmt.Println(" server stoped.")
	})

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
