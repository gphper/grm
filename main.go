package main

import (
	"grm/cmd/run"
	"grm/cmd/user"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "grm"}
	rootCmd.AddCommand(run.CmdRun, user.CmdUser)
	rootCmd.Execute()
}
