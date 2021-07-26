package main

import (
	"httpproj/config"
	"httpproj/db"
	"httpproj/router"
)

func main() {

	//1:初始化配置文件
	config.InitConfig("./config.ini")

	//2：初始化数据库
	db.ConnectDB()
	defer db.CloseDB()

	//3：开启网路监听
	router.InitRouter()

}
