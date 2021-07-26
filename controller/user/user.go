package user

import (
	"encoding/json"
	"fmt"
	"httpproj/model"
	"httpproj/tools"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	//取参数
	data, _ := ioutil.ReadAll(c.Request.Body)
	params := make(map[string]interface{})
	err := json.Unmarshal(data, &params)
	if err != nil {
		tools.JsonMessage(c, 201, "Json格式错误")
		return
	}
	//查数据库
	name, e1 := params["Name"].(string)
	pwd, e2 := params["Password"].(string)
	if e1 == false || e2 == false {
		tools.JsonMessage(c, 201, "字段错误")
		return
	}

	if model.GetUserLoginCount(name, pwd) == 0 {
		tools.JsonMessage(c, 203, "用户名或者密码错误")
	} else {
		//从数据库根据name获取信息，然后保存ID
		user := model.GetUserByName(name)
		tools.JsonObject(c, 200, user)
	}

}

func UserRegister(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	params := make(map[string]interface{})
	err := json.Unmarshal(data, &params)
	if err != nil {
		tools.JsonMessage(c, 201, "Json格式错误")
		return
	}
	fmt.Printf("%s,%s,%s", params["Name"], params["Password"], params["Identifier"])
	fmt.Printf("ctx.Request.body: %v", string(data))
	if model.GetUserCount(params["Name"].(string)) == 0 {
		p := model.User{
			Name:        params["Name"].(string),
			Password:    params["Password"].(string),
			Identifier:  params["Identifier"].(string),
			CreatedTime: time.Now(),
		}
		if p.Add() {
			tools.JsonMessage(c, 1, "注册成功")
		} else {
			tools.JsonMessage(c, tools.DATABASE_ERROR, "上传失败")
		}
	} else {
		tools.JsonMessage(c, 202, "用户已经存在")
	}
}
