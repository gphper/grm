package controllers

import (
	"grm/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type connController struct {
	BaseController
}

var Cc = connController{}

// 获取连接列表
func (con connController) List(c *gin.Context) {
	c.JSON(http.StatusOK, global.RedisServiceStorage)
}
