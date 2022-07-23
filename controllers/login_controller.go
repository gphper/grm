package controllers

import (
	"grm/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type loginController struct {
	BaseController
}

var Loginc = loginController{}

type loginReq struct {
	Username string `form:"username" label:"username" json:"username" binding:"required"`
	Password string `form:"password" label:"password" json:"password" binding:"required"`
}

func (con loginController) Login(c *gin.Context) {
	var loginReq loginReq

	err := con.FormBind(c, &loginReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	password, ok := global.GlobalConf.Accounts[loginReq.Username]
	if !ok || password != loginReq.Password {
		con.Error(c, "账号密码错误")
		return
	}

	key := []byte("TyPbWNRjho")
	// 创建Token结构体
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   loginReq.Username,
		"pass":   loginReq.Password,
		"expire": time.Now().Local().Add(2 * 3600 * time.Second).Format("2006-01-02 15:04:05"),
	})
	// 调用加密方法，发挥Token字符串
	signingString, err := claims.SignedString(key)
	if err != nil {
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"jwt":      signingString,
		"username": loginReq.Username,
	})
}
