package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grm/common"
	"grm/global"

	"github.com/spf13/cobra"
)

var (
	username string
	password string
)

var cmdUserAdd = &cobra.Command{
	Use:   "add",
	Short: "add user",
	Run:   userAddFunc,
}

func userAddFunc(cmd *cobra.Command, args []string) {

	fmt.Println("Please input username:")
	fmt.Scan(&username)
	fmt.Println("Please input password:")
	fmt.Scan(&password)

	fmt.Printf("%s %s", username, password)

	user := global.GlobalConf.Accounts
	user[username] = password
	global.GlobalConf.Accounts = user

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(global.GlobalConf)
	if err != nil {
		fmt.Println(err)
	}
	common.WriteData(buffer.Bytes())
}
