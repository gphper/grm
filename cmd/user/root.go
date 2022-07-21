package user

import "github.com/spf13/cobra"

var CmdUser = &cobra.Command{
	Use:   "user",
	Short: "Option user",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CmdUser.AddCommand(cmdUserAdd)
}
