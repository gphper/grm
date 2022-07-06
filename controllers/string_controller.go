package controllers

import (
	"context"
	"grm/global"
	"grm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type stringController struct {
	BaseController
}

var Sc = stringController{}

// 展示string类型
func (con stringController) Show(c *gin.Context) {

	var req model.KeyReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.RedisServiceStorage[req.Sk].Client

	err = client.Do(context.Background(), "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	data, err := client.Get(context.Background(), req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data": data,
	})
}
