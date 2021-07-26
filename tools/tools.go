package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//返回信息，包含了状态码和对应的信息
func JsonMessage(c *gin.Context, code int, mes string) {
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"detail": mes,
	})
}

//返回Json对象，包含了状态码以及具体需要的对象
func JsonObject(c *gin.Context, code uint, obj interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"obj":    obj,
	})
}
