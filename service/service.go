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
	"grm/glog"
	"net"
	"time"

	"github.com/go-redis/redis"
)

// 生成redis客户端
func NewRedisClient(conf global.RedisService) (*redis.Client, error) {

	ctx := context.Background()
	optConf := &redis.Options{
		Addr:        conf.Config.Addr,
		Password:    conf.Config.Password,
		DialTimeout: 5 * time.Second,
	}
	if conf.UseSsh {
		cli, err := common.GetSSHClient(conf.SSHConfig.SshUsername, conf.SSHConfig.SshPassword, net.JoinHostPort(conf.SSHConfig.SshHost, conf.SSHConfig.SshPort))
		if nil != err {
			return nil, err
		}
		optConf.DialTimeout = -1
		optConf.WriteTimeout = -1
		optConf.ReadTimeout = -1
		optConf.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return cli.Dial(network, addr)
		}
	}

	client := redis.NewClient(optConf)

	client.AddHook(common.RedisLog{
		Logger: glog.NewLogger(conf.RedisService + "/redis.log"),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
