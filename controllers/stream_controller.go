package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"grm/global"
	"grm/model"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type streamController struct {
	BaseController
}

var Stc = streamController{}

// 展示hash类型
func (con streamController) Show(c *gin.Context) {

	type StreamNode struct {
		Num   int    `json:"num"`
		Id    string `json:"id"`
		Value string `json:"value"`
	}

	var first string
	var last string
	var xSlice []redis.XMessage
	var req model.StreamKeyReq
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

	total, err := client.XLen(context.Background(), req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	total_page := math.Ceil(float64(total) / float64(req.Limit))
	if req.Page == int(total_page) {
		req.Limit = int(total) - (int(total_page)-1)*req.Limit
	}

	data := make([]StreamNode, req.Limit)

	limit := req.Limit + 1
	if req.Type == "next" {
		xSlice, err = client.XRangeN(context.Background(), req.Id, req.Item, "+", int64(limit)).Result()
		if err != nil {
			con.Error(c, err.Error())
			return
		}

		tmp := xSlice[len(xSlice)-1]
		last = tmp.ID
		first = req.Item
		if req.Page != int(total_page) {
			xSlice = xSlice[:len(xSlice)-1]
		}

	} else {
		xSlice, err = client.XRevRangeN(context.Background(), req.Id, req.Item, "-", int64(limit)).Result()
		if err != nil {
			con.Error(c, err.Error())
			return
		}
		fmt.Println(req.Item)
		fmt.Println(xSlice)
		tmp := xSlice[len(xSlice)-1]
		first = tmp.ID
		last = req.Item
		xSlice = xSlice[1:]
	}

	start := req.Limit * (req.Page - 1)
	for k, v := range xSlice {
		tmp, _ := json.Marshal(v.Values)

		data[k] = StreamNode{
			Num:   start + k + 1,
			Id:    v.ID,
			Value: string(tmp),
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
		"first": first,
		"last":  last,
	})
}

// 删除指定数据
func (con streamController) Del(c *gin.Context) {
	var req model.ZsetItemKeyReq

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

	count, err := client.XDel(context.Background(), req.Id, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}

// 添加item
func (con streamController) AddItem(c *gin.Context) {
	var req model.AddStreamItemKeyReq

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

	var item map[string]interface{}
	err = json.Unmarshal([]byte(req.Item), &item)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	result, err := client.XAdd(context.Background(), &redis.XAddArgs{
		Stream:       req.Id,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           req.Idx,
		Values:       item,
	}).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data": result,
	})
}
