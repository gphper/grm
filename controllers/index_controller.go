package controllers

import (
	"context"
	"fmt"
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
		DbName string `json:"db"`
		DbNum  string `json:"keys"`
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
			DbName: key,
			DbNum:  num,
		}
	}

	con.Success(c, http.StatusOK, dbInfos)
}
