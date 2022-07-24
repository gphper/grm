package controllers

import (
	"context"
	"grm/global"
	"grm/model"
	"math"
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

	start := (req.Page - 1) * req.Limit
	end := start + req.Limit - 1
	if end > int(total-1) {
		end = int(total - 1)
	}
	data := make([]ListNode, end-start+1)

	setSlice, _, err := client.SScan(ctx, req.Id, uint64(start), "*", int64(end-start+1)).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	for k, v := range setSlice {
		data[k] = ListNode{
			Num:   start + k + 1,
			Value: v,
		}
	}

	ttl, err := client.TTL(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data":  data,
		"page":  req.Page,
		"total": math.Ceil(float64(total) / float64(req.Limit)),
		"ttl":   ttl.Seconds(),
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
