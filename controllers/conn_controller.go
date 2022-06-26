package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"grm/global"
	"grm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type connController struct {
	BaseController
}

var Cc = connController{}

// 获取连接列表
func (con connController) List(c *gin.Context) {
	type ConnItem struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
	conns := make([]ConnItem, 0)
	for k, v := range global.RedisServiceStorage {
		conns = append(conns, ConnItem{
			Key:  k,
			Name: v.RedisService,
		})
	}
	con.Success(c, http.StatusOK, conns)
}

// 保存连接信息
func (con connController) Add(c *gin.Context) {
	var conf model.ServiceConfigReq
	err := con.FormBind(c, &conf)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	m := md5.New()
	m.Write([]byte(conf.ServiceName))
	key := hex.EncodeToString(m.Sum(nil))

	global.RedisServiceStorage[key] = global.RedisService{
		RedisService: conf.ServiceName,
		UseSsh:       conf.UseSsh,
		SSHConfig: model.SSHConfig{
			SshHost:     conf.SshHost,
			SshPort:     conf.SshPort,
			SshUsername: conf.SshUsername,
			SshPassword: conf.SshPassword,
		},
	}

	con.Success(c, http.StatusOK, gin.H{
		"key":  key,
		"name": conf.ServiceName,
	})
}
