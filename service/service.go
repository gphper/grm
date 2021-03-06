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
	"net"

	"github.com/go-redis/redis"
)

// 生成redis客户端
func NewRedisClient(conf global.RedisService) (*redis.Client, error) {

	ctx := context.Background()
	optConf := &redis.Options{}
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
		Logger: common.NewLogger(conf.RedisService),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
