package controllers

import (
	"context"
	"grm/global"
	"grm/model"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type hashController struct {
	BaseController
}

var Hc = hashController{}

// 展示hash类型
func (con hashController) Show(c *gin.Context) {

	type HashNode struct {
		Num   int    `json:"num"`
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var req model.HashKeyReq
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

	total, err := client.HLen(context.Background(), req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	start := (req.Page - 1) * req.Limit
	end := start + req.Limit - 1
	if end > int(total-1) {
		end = int(total - 1)
	}
	data := make([]HashNode, 0)

	hashSlice, _, err := client.HScan(context.Background(), req.Id, 0, "*", int64(end-start+1)).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	len := len(hashSlice)
	tmp := HashNode{}
	num := 1
	for i := 0; i < len; i++ {
		if i%2 == 0 {
			tmp.Key = hashSlice[i]
			tmp.Num = num
		} else {
			tmp.Value = hashSlice[i]
			data = append(data, tmp)
			tmp = HashNode{}
			num++
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
func (con hashController) Del(c *gin.Context) {
	var req model.HashItemKeyReq

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

	count, err := client.HDel(context.Background(), req.Id, req.Item).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}

// 添加item
func (con hashController) AddItem(c *gin.Context) {
	var req model.AddHashItemKeyReq

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

	count, err := client.HSet(context.Background(), req.Id, req.Itemk, req.Itemv).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}
