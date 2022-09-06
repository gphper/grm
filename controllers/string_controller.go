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

	client := global.GetClient(req.Sk)
	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	data, err := client.Get(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	ttl, err := client.TTL(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data": data,
		"ttl":  ttl.Seconds(),
	})
}

// 添加
func (con stringController) Add(c *gin.Context) {
	var req model.AddStrReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GetClient(req.Sk)
	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	res, err := client.Set(ctx, req.Id, req.Item, 0).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data": res,
	})
}
