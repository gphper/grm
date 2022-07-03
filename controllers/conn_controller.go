package controllers

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"grm/global"
	"grm/model"
	"grm/service"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
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
		Config: &redis.Options{
			Addr:     net.JoinHostPort(conf.Host, conf.Port),
			Password: conf.Password,
			DB:       0,
		},
		UseSsh: conf.UseSsh,
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

// 测试连接
func (con connController) TestConn(c *gin.Context) {
	var req model.ServiceConfigReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	// optionConfig := &redis.Options{
	// 	Addr:     net.JoinHostPort(conf.Host, conf.Port),
	// 	Password: conf.Password,
	// 	DB:       0,
	// }

	// if conf.UseSsh {
	// 	optionConfig.DialTimeout = -1
	// 	optionConfig.WriteTimeout = -1
	// 	optionConfig.ReadTimeout = -1
	// 	cli, err := common.GetSSHClient(conf.SshUsername, conf.SshPassword, conf.SshHost+":"+conf.SshPort)
	// 	if nil != err {
	// 		panic(err)
	// 	}
	// 	optionConfig.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
	// 		return cli.Dial(network, addr)
	// 	}
	// } else {
	// 	optionConfig.DialTimeout = 5 * time.Second
	// }

	// client := redis.NewClient(optionConfig)

	client, err := service.NewRedisClient(global.RedisService{
		RedisService: req.ServiceName,
		Config: &redis.Options{
			Addr:     net.JoinHostPort(req.Host, req.Port),
			Password: req.Password,
			DB:       0,
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
