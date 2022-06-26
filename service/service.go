/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 14:45:08
 */
package service

import (
	"context"
	"grm/common"
	"grm/global"
	"grm/model"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/ssh"
)

func AddServiceConf(conf model.ServiceConfigReq) (err error) {
	var cli *ssh.Client

	optionConfig := &redis.Options{
		Addr:     conf.Host + ":" + string(conf.Port),
		Password: conf.Password,
		DB:       0,
	}

	ctx := context.Background()
	if conf.UseSsh {
		cli, err = common.GetSSHClient(conf.SSHConfig.SshUsername, conf.SSHConfig.SshPassword, conf.SSHConfig.SshHost+":"+string(conf.SSHConfig.SshPort))
		if nil != err {
			return
		}
		optionConfig.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return cli.Dial(network, addr)
		}
	}

	client := redis.NewClient(optionConfig)

	client.AddHook(common.RedisLog{
		Logger: common.NewLogger(conf.ServiceName),
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return
	}

	RsSlice := global.RedisService{
		RedisService: conf.ServiceName,
		Config:       optionConfig,
		UseSsh:       conf.UseSsh,
		SSHConfig:    conf.SSHConfig,
	}

	global.RedisServiceStorage[conf.ServiceName] = RsSlice

	//设置全局参数
	global.UseClient.ConnectName = conf.ServiceName
	global.UseClient.Db = 0
	global.UseClient.Client = client

	return nil
}

func ServiceSwitch(conf model.ServiceSwitchReq) (err error) {
	config := global.RedisServiceStorage[conf.Service]
	config.Config.DB = conf.Db

	client := redis.NewClient(config.Config)
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	global.UseClient.ConnectName = conf.Service
	global.UseClient.Db = conf.Db
	global.UseClient.Client = redis.NewClient(config.Config)

	return nil
}

func SearchKeyType(conf model.RedisKeyReq, c *gin.Context) (keys []string, cursor uint64, err error) {

	val, _ := c.Get("username")
	ctx := context.WithValue(context.Background(), "username", val)

	if conf.SearchType == 1 {
		//查询指定值
		keys = append(keys, conf.SearchKey)

		_, err = global.UseClient.Client.Exists(ctx, conf.SearchKey).Result()
		return
	} else {
		//模糊匹配
		keys, cursor, err = global.UseClient.Client.Scan(ctx, conf.Course, conf.SearchKey, global.Limit).Result()
	}
	return
}
