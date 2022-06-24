/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 10:43:45
 */
package common

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func GetSSHClient(user, pass, addr string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		//需要验证服务端，不做验证返回nil，没有该参数会报错
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	sshConn, err := ssh.Dial("tcp", addr, config)
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("world")
	return sshConn, nil
}
