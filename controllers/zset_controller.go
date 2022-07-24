package controllers

import (
	"context"
	"grm/global"
	"grm/model"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type zsetController struct {
	BaseController
}

var Zc = zsetController{}

// 展示hash类型
func (con zsetController) Show(c *gin.Context) {

	type ZsetNode struct {
		Num   int    `json:"num"`
		Score int    `json:"score"`
		Value string `json:"value"`
	}

	var req model.ZsetKeyReq
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

	total, err := client.ZCard(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	start := (req.Page - 1) * req.Limit
	end := start + req.Limit - 1
	if end > int(total-1) {
		end = int(total - 1)
	}

	zsetSlice, err := client.ZRangeWithScores(ctx, req.Id, int64(start), int64(end)).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	data := make([]ZsetNode, len(zsetSlice))
	for k, v := range zsetSlice {
		data[k] = ZsetNode{
			Num:   start + k + 1,
			Score: int(v.Score),
			Value: v.Member.(string),
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
func (con zsetController) Del(c *gin.Context) {
	var req model.ZsetItemKeyReq

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

	count, err := client.ZRem(ctx, req.Id, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}

// 添加item
func (con zsetController) AddItem(c *gin.Context) {
	var req model.AddZsetItemKeyReq

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

	score, _ := strconv.Atoi(req.Score)
	count, err := client.ZAdd(ctx, req.Id, &redis.Z{
		Score:  float64(score),
		Member: req.Item,
	}).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}
