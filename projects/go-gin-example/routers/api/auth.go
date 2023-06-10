package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/fcf/go-gin-example/models"
	"github.com/fcf/go-gin-example/pkg/e"
	"github.com/fcf/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist, err := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else if err != nil {
			fmt.Printf("auth is not exist: %x", err)
			code = e.ERROR_AUTH
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func RegistAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.MaxSize(username, 50, "username").Message("用户名最长50字节")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MaxSize(password, 50, "password").Message("密码最长50字节")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		exist, _ := models.ExistAuthByName(username)
		if !exist {
			fmt.Printf("xxxxxxxx fcf %s 未注册过", username)
			suc := models.RegistAuth(username, password)
			if suc {
				fmt.Println("regist auth success")
				code = e.SUCCESS
			} else {
				fmt.Println("regist auth failed")
				code = e.ERROR
			}
		} else {
			code = e.ERROR_EXIST_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
