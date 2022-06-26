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
	type ConnItem struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
	conns := make([]ConnItem, 0)
	for k, v := range global.RedisServiceStorage {
		conns = append(conns, ConnItem{
			Key:  k,
			Name: v.RedisService,
		})
	}
	con.Success(c, http.StatusOK, conns)
}
