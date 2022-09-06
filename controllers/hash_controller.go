package controllers

import (
	"context"
	"grm/global"
	"grm/model"
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

	client := global.GetClient(req.Sk)

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	total, err := client.HLen(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	hashSlice, cursor, err := client.HScan(ctx, req.Id, uint64(req.Cursor), "*", 100).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	data := make([]HashNode, 0)
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

	ttl, err := client.TTL(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"data":   data,
		"total":  total,
		"ttl":    ttl.Seconds(),
		"cursor": cursor,
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

	client := global.GetClient(req.Sk)

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	count, err := client.HDel(ctx, req.Id, req.Item).Result()
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

	client := global.GetClient(req.Sk)

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", req.Db).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	count, err := client.HSet(ctx, req.Id, req.Itemk, req.Itemv).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"count": count,
	})
}
