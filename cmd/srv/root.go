package srv

import "github.com/spf13/cobra"

var CmdSrv = &cobra.Command{
	Use:   "srv",
	Short: "Option server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CmdSrv.AddCommand(cmdSrvRun)
}
