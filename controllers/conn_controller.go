package controllers

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"grm/common"
	"grm/global"
	"grm/model"
	"grm/service"
	"net"
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

	for k, v := range global.GlobalConf.RedisServices {
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

	global.GlobalConf.RedisServices[key] = global.RedisService{
		RedisService: conf.ServiceName,
		Config: model.RedisConfig{
			Addr:     net.JoinHostPort(conf.Host, conf.Port),
			Password: conf.Password,
			Db:       0,
		},
		UseSsh: conf.UseSsh,
		SSHConfig: model.SSHConfig{
			SshHost:     conf.SshHost,
			SshPort:     conf.SshPort,
			SshUsername: conf.SshUsername,
			SshPassword: conf.SshPassword,
		},
	}

	//存储
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err = encoder.Encode(global.GlobalConf)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	common.WriteData(buffer.Bytes())

	con.Success(c, http.StatusOK, gin.H{
		"key":  key,
		"name": conf.ServiceName,
	})
}

// 删除连接信息
func (con connController) Del(c *gin.Context) {
	var req model.DelServiceReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	delete(global.GlobalConf.RedisServices, req.ServiceName)

	//存储
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err = encoder.Encode(global.GlobalConf)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	common.WriteData(buffer.Bytes())

	con.Success(c, http.StatusOK, gin.H{
		"data": req.ServiceName,
	})
}

// 测试连接
func (con connController) TestConn(c *gin.Context) {
	var req model.ServiceConfigReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	client, err := service.NewRedisClient(global.RedisService{
		RedisService: req.ServiceName,
		Config: model.RedisConfig{
			Addr:     net.JoinHostPort(req.Host, req.Port),
			Password: req.Password,
			Db:       0,
		},
		UseSsh: req.UseSsh,
		SSHConfig: model.SSHConfig{
			SshHost:     req.SshHost,
			SshPort:     req.SshPort,
			SshUsername: req.SshUsername,
			SshPassword: req.SshPassword,
		},
	})

	if err != nil {
		con.Error(c, err.Error())
		return
	}
	defer client.Close()

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, http.StatusOK, "")
}
