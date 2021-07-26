package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
)

var Config *goconfig.ConfigFile

func InitConfig(path string) {
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		log.Fatal("读取配置文件失败[config.ini]")
		return
	}
	Config = cfg
}
func GetValue(sec string, key string) string {
	value, err := Config.GetValue(sec, key)
	fmt.Println(value)
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", value, err)
	}
	return value
}

func GetPort() string {
	value, err := Config.GetValue("dev", "port")
	fmt.Println(value)
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", value, err)
	}
	return value
}
func GetFileServer() string {
	value, err := Config.GetValue("file", "FileServer")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", value, err)
	}
	return value
}

func GetLocalPath() string {
	value, err := Config.GetValue("file", "LocalPath")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", value, err)
	}
	return value
}

type MySqlSetting struct {
	UserName string
	Password string
	IP       string
	DataName string
}

func GetSqlSetting() *MySqlSetting {
	sec, _ := Config.GetSection("mysql")
	myconfig := &MySqlSetting{}
	myconfig.UserName = sec["name"]
	myconfig.Password = sec["password"]
	myconfig.IP = sec["ip"]
	myconfig.DataName = sec["database"]
	return myconfig
}

func GetDataBaseConfig() map[string]string {
	sec, err := Config.GetSection("mysql")
	if err != nil {
		log.Fatal("获取参数失败")
	}
	mapConfig := make(map[string]string)
	mapConfig["userName"] = sec["name"]
	mapConfig["password"] = sec["password"]
	mapConfig["ip"] = sec["ip"]
	mapConfig["database"] = sec["database"]
	return mapConfig
}
