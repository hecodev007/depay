package config

import (
	"github.com/BurntSushi/toml"
	//"go-gin-api/tools/store"
)

var GlobalConf *Config

type Conf struct {
	HostPort string
	Username string
	DBName   string
	Password string
	MaxConns int
	MaxIdle  int
	ShowSQL  bool
}

type Config struct {
	//Runtime string //release/debug
	Name  string
	Port  string
	Mysql Conf
}

func New() *Config {
	return &Config{}
}

func (p *Config) Init(cfgFile string) error {
	_, err := toml.DecodeFile(cfgFile, p)
	GlobalConf = p
	return err
}
