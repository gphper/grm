/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 10:43:45
 */
package common

import (
	"fmt"
	"net"
	"time"

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
		Timeout: 10 * time.Second,
	}

	sshConn, err := net.Dial("tcp", addr)
	if nil != err {
		fmt.Println("net dial err: ", err)
		return nil, err
	}

	clientConn, chans, reqs, err := ssh.NewClientConn(sshConn, addr, config)
	if nil != err {
		sshConn.Close()
		fmt.Println("ssh client conn err: ", err)
		return nil, err
	}

	client := ssh.NewClient(clientConn, chans, reqs)

	return client, nil
}
