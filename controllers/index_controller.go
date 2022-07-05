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
	redisServer := global.RedisServiceStorage[keys[1]]

	client, err := service.NewRedisClient(redisServer)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	redisServer.Client = client
	global.RedisServiceStorage[keys[1]] = redisServer

	result, err := client.Info(context.Background(), "Keyspace").Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	dbNum, err := client.ConfigGet(context.Background(), "databases").Result()
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

	maxDb, _ := strconv.Atoi(dbNum["databases"])

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
	client := global.RedisServiceStorage[dbInfo[1]].Client

	err = client.Do(context.Background(), "select", index).Err()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	keys, _, err := client.Scan(context.Background(), 0, "*", 10000).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	gen := common.NewTrie()

	for _, v := range keys {
		stringSlice := strings.Split(v, ":")
		gen.Insert(stringSlice, v)
	}

	con.Success(c, http.StatusOK, common.GetOne(gen.Root.Children, "", dbInfo[1], index))
}

// 查看key值详情
func (con indexController) GetKeyType(c *gin.Context) {
	var req model.ShowKeyReq
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

	types, err := client.Type(context.Background(), req.Id).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, gin.H{
		"types": types,
	})
}
