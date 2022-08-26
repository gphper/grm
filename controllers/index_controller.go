package controllers

import (
	"bytes"
	"context"
	"encoding/json"
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
	var err error
	var cursor uint64

	result := make(map[string]interface{})
	keys := make([]string, 0)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	cursor = uint64(req.Cursor)

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

	for {
		var tmp []string
		tmp, cursor, err = client.Scan(ctx, cursor, req.Match, 10000).Result()
		if err != nil {
			con.Error(c, err.Error())
			return
		}

		keys = append(keys, tmp...)

		if len(keys) > 9000 || cursor == 0 {
			break
		}
	}

	gen := common.NewTrie()

	for _, v := range keys {
		stringSlice := make([]string, 0)
		if global.GlobalConf.Tree {
			stringSlice = strings.Split(v, global.GlobalConf.Separator)
		} else {
			stringSlice = append(stringSlice, v)
		}
		gen.Insert(stringSlice, v)
	}

	commonNode := common.GetOne(gen.Root.Children, "", dbInfo[1], index)
	if cursor > 0 {
		commonNode = append(commonNode, common.Node{
			Title:      "",
			All:        "nextpagegrmtag",
			ServiceKey: "",
			Db:         0,
			Children:   []common.Node{},
		})
	}

	result["data"] = commonNode
	result["count"] = len(keys)
	result["cursor"] = cursor
	con.Success(c, http.StatusOK, result)
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
	var req model.DelKeysReq
	keys := make([]string, 0)

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

	if req.Single > 0 {
		keys, _, err = client.Scan(ctx, 0, req.Id+"*", 10000).Result()
		if err != nil {
			con.Error(c, err.Error())
			return
		}
	} else {
		keys = append(keys, req.Id)
	}

	_, err = client.Del(ctx, keys...).Result()
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

// 系统设置
func (con indexController) Setting(c *gin.Context) {

	if c.Request.Method == "GET" {
		con.Success(c, http.StatusOK, gin.H{
			"tree":      global.GlobalConf.Tree,
			"separator": global.GlobalConf.Separator,
		})
	} else {
		var req model.SettingReq
		err := con.FormBind(c, &req)
		if err != nil {
			con.Error(c, err.Error())
			return
		}

		global.GlobalConf.Separator = req.Separator
		global.GlobalConf.Tree = req.Tree

		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		err = encoder.Encode(global.GlobalConf)
		if err != nil {
			con.Error(c, err.Error())
			return
		}

		common.WriteData(buffer.Bytes())

		con.Success(c, http.StatusOK, gin.H{
			"data": "",
		})
	}

}
