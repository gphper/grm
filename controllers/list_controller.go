package controllers

import (
	"context"
	"grm/global"
	"grm/model"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listController struct {
	BaseController
}

var Lc = listController{}

// 展示string类型
func (con listController) Show(c *gin.Context) {

	type ListNode struct {
		Num   int    `json:"num"`
		Value string `json:"value"`
	}

	var req model.ListKeyReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GlobalClients[req.Sk]

	err = client.Do(context.Background(), "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	total, err := client.LLen(context.Background(), req.Id).Result()
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

	listSlice, err := client.LRange(context.Background(), req.Id, int64(start), int64(end)).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	for k, v := range listSlice {
		data[k] = ListNode{
			Num:   start + k + 1,
			Value: v,
		}
	}

	ttl, err := client.TTL(context.Background(), req.Id).Result()
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
func (con listController) Del(c *gin.Context) {
	var req model.ListItemKeyReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GlobalClients[req.Sk]

	err = client.Do(context.Background(), "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	count, err := client.LRem(context.Background(), req.Id, 1, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}

// 添加item
func (con listController) AddItem(c *gin.Context) {
	var req model.ListItemKeyReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client := global.GlobalClients[req.Sk]

	err = client.Do(context.Background(), "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	count, err := client.LPush(context.Background(), req.Id, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}
