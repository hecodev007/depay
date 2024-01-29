package main

import (
	"depay/contract"
	"depay/middleware"
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
	//0 本地
	//1 生产
	env := flag.Int64("e", 0, "owner private")
	flag.Parse()

	loger.InitLog(".", "reward", "")
	cfgPath := "./config/config.toml"
	if *env == 1 {
		cfgPath = "./config/config_pro.toml"
	}

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
	r.Use(middleware.Cors())
	r.Use(middleware.SetUp())
	router := r.Group("api")
	router.POST("/genPayOrder", s.GenPayOrder)
	router.GET("/getPayOrder", s.GetPayOrder)
	router.GET("/getOrderStatus", s.GetOrderStatus)
	router.POST("/cancelPayOrder", s.CancelPayOrder)
	router.GET("/testApi", s.TestApi)
	//r.OPTIONS("/testApi", s.TestApi)
	fmt.Println("start serve。。。")
	go func() {
		for {
			select {
			//bsc: '0x97E1614db4E8f1Db69d08F0eBb252749E2085f15',
			//	goerli: '0xD28bbAD2290F87CEDe36Db234a39C984F6DD4124',
			//	maticmum: '0x888599BA44f703699F6E27BF6187633637aF16Cd',
			case <-time.After(5 * time.Second):
				contract.FilOne("Bsc", "https://bsc-pokt.nodies.app", []string{"0x97E1614db4E8f1Db69d08F0eBb252749E2085f15"})
				contract.FilOne("PolygonTest", "https://polygon-mumbai-bor.publicnode.com", []string{"0x888599BA44f703699F6E27BF6187633637aF16Cd"})
				contract.FilOne("Goerli", "https://ethereum-goerli.publicnode.com", []string{"0xD28bbAD2290F87CEDe36Db234a39C984F6DD4124"})

			}
			//insert block_height(chain,height) values("Goerli",10435515);

		}
	}()
	go func() {
		for {
			select {
			case <-time.After(15 * time.Second):
				service.Notify()
			}

		}
	}()
	r.Run(conf.Port)
	// 捕捉退出信号
	d := death.NewDeath(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL,
		syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGALRM)
	d.WaitForDeathWithFunc(func() {
		fmt.Println(" server stoped.")
	})

}

//priv: MIICXQIBAAKBgQDhhCkKiF+CoxwohxlM6R39AxbVWpwn7aqNOeRDXm21wZss/2OL/jWbHm1VHqvyjDOcYoO59B5MbEwZdx7H6BFS3l1aG1b9WE6C4t4FEgIZBsxa0u4RNsxa+nCt2YMsK3nDJjBMTuaFSLgueaWnhKK0xM9b+j069BBLcWXoCMsnJQIDAQABAoGALJs6jvXIhouCz1VhoL/fiaMpygvBJXiyKnsP9m9gHfpsmirt4svmiIctMw/9DN3Ee6NU0NxDffRR3RudwAbcHfwGVV68jXguHpAJsd00kz7orZlLHbPwDQnn83LVF7cgHDTwVZWc1mpzGh6UZpnwR3l04vhePOkYSoRT3y3jm6ECQQDwcAyU7fhAAcPHVq54RCzE2VgHWsIX0qr1TuZTacVwzTS6KLcqRcujI+022poKays9Ej8UpS5+Aq+9IJTPJcuJAkEA8BzehGrLDy8g+mK06+MUp0zWPj2t+NiHwPDGvPg7kwrPN0VJT1DXhUeyzSkmlz/nBmD6Rc3580E4AF+4Br0LvQJBALW/vYMGr9WSf++7Mn9u6XiT4tsMXBOuB9UPI0SCe+Fc/TKLfInT4K8dhT8l17Nwd2re1BhDFPXkCfwpGNPNeiECQGXf+dE49lrE5jsV8ik7OaIaCbRyuwOf60lDXy8CK1Sh+3U54nbSl/6mgwhk80itBjpAny9Ky0gYXcha1FuXjgkCQQDUPxa8zSPyj9ygaON5R/0YdqcLAhzu3hS+uJSRq2VIvTG+lFW42VZ+snFnTZqkM7rzPx9EVM0pCvWjCY9be1CI
//public: MIGJAoGBAOGEKQqIX4KjHCiHGUzpHf0DFtVanCftqo055ENebbXBmyz/Y4v+NZsebVUeq/KMM5xig7n0HkxsTBl3HsfoEVLeXVobVv1YToLi3gUSAhkGzFrS7hE2zFr6cK3ZgywrecMmMExO5oVIuC55paeEorTEz1v6PTr0EEtxZegIyyclAgMBAAE=

// timeStamp
func Router() *gin.Engine {

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
		return nil
	} else {
		model.DB = db
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s := service.Service{
		Node: node,
	}
	r := gin.Default()
	//r.GET("/testApi", s.TestApi)
	r.GET("/testApi", s.TestApi)
	//r.Any("/testApi", s.TestApi)
	r.POST("/genPayOrder", s.GenPayOrder)
	r.GET("/getPayOrder", s.GetPayOrder)
	r.GET("/getOrderStatus", s.GetOrderStatus)
	r.POST("/cancelPayOrder", s.CancelPayOrder)
	//	r.POST("/addMerchant", s.AddMerchant)

	//fmt.Println("start serve。。。")

	//r.Run(conf.Port)

	return r
}
