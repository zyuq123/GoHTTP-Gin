package router

import (
	"httpproj/config"
	"httpproj/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.Default()
	//1:登录
	r.POST("/login", user.UserLogin)
	//2：注册
	r.POST("/register", user.UserRegister)

	r.Run(":" + config.GetPort())
}
