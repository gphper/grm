package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grm/common"
	"grm/global"

	"github.com/spf13/cobra"
)

var delusername string

var cmdUserDel = &cobra.Command{
	Use:   "delete",
	Short: "delete user",
	Run:   userDelFunc,
}

func userDelFunc(cmd *cobra.Command, args []string) {

	fmt.Println("Please input username:")
	fmt.Scan(&delusername)

	user := global.GlobalConf.Accounts
	delete(user, delusername)

	global.GlobalConf.Accounts = user

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(global.GlobalConf)
	if err != nil {
		fmt.Println(err)
	}
	common.WriteData(buffer.Bytes())

	fmt.Println("delete success!")
}
