package controllers

import (
	"context"
	"fmt"
	"grm/common"
	"grm/global"
	"grm/model"
	"grm/service"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type indexController struct {
	BaseController
}

var Ic = indexController{}

// 获取连接列表
func (con indexController) Open(c *gin.Context) {

	type dbInfo struct {
		DbName     string `json:"db"`
		DbNum      string `json:"keys"`
		ServiceKey string `json:"servicekey"`
	}

	var req model.OpenDbReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	keys := strings.Split(req.Index, "_")
	redisServer := global.GlobalConf.RedisServices[keys[1]]
	client, err := service.NewRedisClient(redisServer)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	global.GlobalClients[keys[1]] = client
	global.GlobalConf.RedisServices[keys[1]] = redisServer

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	result, err := client.Info(ctx, "Keyspace").Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	dbNum, err := client.ConfigGet(ctx, "databases").Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	Slice := strings.Split(result, "\r\n")

	dbs := make(map[string]string)
	for _, dbInfo := range Slice[1:] {
		compileNumRegex := regexp.MustCompile(`(.*?):keys=(\d+)`)
		keyNum := compileNumRegex.FindSubmatch([]byte(dbInfo))

		if len(keyNum) != 0 {
			dbs[string(keyNum[1])] = string(keyNum[2])
		}
	}

	databases := dbNum[1].(string)

	maxDb, err := strconv.Atoi(databases)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	dbInfos := make([]dbInfo, maxDb)
	for i := 0; i < maxDb; i++ {
		key := fmt.Sprintf("db%d", i)

		num := "0"
		if numv, ok := dbs[key]; ok {
			num = numv
		}

		dbInfos[i] = dbInfo{
			DbName:     key,
			DbNum:      num,
			ServiceKey: keys[1],
		}
	}

	con.Success(c, http.StatusOK, dbInfos)
}

// 获取所有keys
func (con indexController) GetKeys(c *gin.Context) {
	var req model.GetKeysReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	dbInfo := strings.Split(req.Index, "-")
	index, _ := strconv.Atoi(dbInfo[0])
	client := global.GlobalClients[dbInfo[1]]

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	err = client.Do(ctx, "select", index).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	keys, cursor, err := client.Scan(ctx, uint64(req.Cursor), req.Match, 1000).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	gen := common.NewTrie()

	for _, v := range keys {
		stringSlice := strings.Split(v, ":")
		gen.Insert(stringSlice, v)
	}

	con.Success(c, http.StatusOK, gin.H{
		"data":   common.GetOne(gen.Root.Children, "", dbInfo[1], index),
		"count":  len(keys),
		"cursor": cursor,
	})
}

// 查看key值详情
func (con indexController) GetKeyType(c *gin.Context) {
	var req model.KeyReq
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

	types, err := client.Type(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"types": types,
	})
}

// 删除key值
func (con indexController) DelKey(c *gin.Context) {
	var req model.KeyReq
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

	_, err = client.Del(ctx, req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"status": 1,
	})
}

// TTL 设置
func (con indexController) TtlKey(c *gin.Context) {
	var req model.TtlKeyReq

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

	ttl, err := strconv.Atoi(req.Ttl)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	ok, err := client.Expire(ctx, req.Id, time.Duration(ttl)*time.Second).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"result": ok,
	})
}

func (con indexController) SerInfo(c *gin.Context) {
	var req model.InfoReq

	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	redisServer := global.GlobalConf.RedisServices[req.Key]

	client, err := service.NewRedisClient(redisServer)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	result, err := client.Info(ctx).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"info": result,
		"name": redisServer.RedisService,
	})
}

func (con indexController) LuaRun(c *gin.Context) {
	var req model.LuaRunReq

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

	args := make([]interface{}, len(req.Argv))
	for k, v := range req.Argv {
		args[k] = v
	}

	res, err := client.Eval(ctx, req.Script, req.Keys, args...).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	var resultPut string
	switch val := res.(type) {
	case string:
		resultPut = val
	case int:
		resultPut = strconv.Itoa(val)
	case []interface{}:
		for _, v := range val {
			resultPut += fmt.Sprintf("%s \r\n", v.(string))
		}
	case bool:
		if val {
			resultPut = "True"
		} else {
			resultPut = "False"
		}
	}

	con.Success(c, http.StatusOK, gin.H{
		"data": resultPut,
	})
}
