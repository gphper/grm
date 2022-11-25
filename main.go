package main

import (
	"grm/cmd/run"
	"grm/cmd/srv"
	"grm/cmd/user"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "grm"}
	rootCmd.AddCommand(run.CmdRun, user.CmdUser, srv.CmdSrv)
	rootCmd.Execute()
}
