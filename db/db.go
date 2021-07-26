package db

import (
	"httpproj/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	conf := config.GetDataBaseConfig()
	//db, err = gorm.Open("mysql",
	//	"root:密码@/user?charset=utf8&parseTime=True&loc=Local")
	connStr := conf["userName"] + ":" +
		conf["password"] + "@/" +
		conf["database"] +
		"?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	//测试连接是不是通顺
	if DB.DB().Ping() != nil {
		panic("error ")
	}
	DB.DB().SetMaxIdleConns(50)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(true)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
