package controllers

import (
	"context"
	"grm/global"
	"grm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type setController struct {
	BaseController
}

var Setc = setController{}

// 展示string类型
func (con setController) Show(c *gin.Context) {

	type ListNode struct {
		Num   int    `json:"num"`
		Value string `json:"value"`
	}

	var req model.SetKeyReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GlobalClients[req.Sk]

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	total, err := client.SCard(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	setSlice, cursor, err := client.SScan(ctx, req.Id, uint64(req.Cursor), "*", 100).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	data := make([]ListNode, len(setSlice))

	for k, v := range setSlice {
		data[k] = ListNode{
			Num:   k + 1,
			Value: v,
		}
	}

	ttl, err := client.TTL(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data":   data,
		"cursor": cursor,
		"total":  total,
		"ttl":    ttl.Seconds(),
	})
}

// 删除指定数据
func (con setController) Del(c *gin.Context) {
	var req model.SetItemKeyReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GlobalClients[req.Sk]

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	count, err := client.SRem(ctx, req.Id, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}

// 添加item
func (con setController) AddItem(c *gin.Context) {
	var req model.SetItemKeyReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GlobalClients[req.Sk]

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	count, err := client.SAdd(ctx, req.Id, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}
