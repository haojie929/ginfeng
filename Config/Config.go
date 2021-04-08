package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	//Yaml "gopkg.in/yaml.v2"
)

type ConfigHandler struct {
	SysConfig   System           `yaml:"system"`
	GromConfig  map[string]Mysql `yaml:"grom"`
	RedisConfig map[string]Redis `yaml:"redis"`
}

var BaseConfig *ConfigHandler

// 初始化配置
func init() {
	BaseConfig = &ConfigHandler{
		SysConfig: System{
			AppName:   "gin-feng",
			Env:       "public",
			HTTPAddr:  "127.0.0.1:8080",
			HTTPSAddr: "127.0.0.1:10443",
		},
	}
}

// 初始化
func (cs *ConfigHandler) Init() {
	if cs.SysConfig.Env = os.Getenv("GIN_RUNMODE"); cs.SysConfig.Env == "" {
		cs.SysConfig.Env = "debug"
	}

	configInfo, err := ReadYamlConfig("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	BaseConfig = configInfo
}

func (cs *ConfigHandler) GetMysqlConfig() map[string]Mysql {
	return BaseConfig.GromConfig
}

func (cs *ConfigHandler) GetRedisConfig() map[string]Redis {
	return BaseConfig.RedisConfig
}

//func (cs *ConfigHandler) GetAllMysqlConnect()(map[string]string) {
//	sqlConnects := make(map[string]string)
//	var sqlDb string
//	for i:=0;i<len(cs.SysConfig.DbConnect);i++{
//		sqlDb = cs.SysConfig.DbConnect[i]
//		mysqldb := cs.MysqlConfig[sqlDb]
//		sqlConnects[sqlDb] = mysqldb.Username + ":" + mysqldb.Password + "@tcp(" + mysqldb.Path + ")/" + mysqldb.Dbname + "?" + mysqldb.Config
//	}
//	return sqlConnects
//}

func ReadYamlConfig(path string) (*ConfigHandler, error) {
	BaseConfig := &ConfigHandler{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	//把yaml形式的字符串解析成struct类型
	if yaml.Unmarshal(data, BaseConfig) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	fmt.Println(BaseConfig)
	return BaseConfig, nil
}
