package main

import (
	"github.com/fengchunjian/goexamples/cobra/node"
	"github.com/spf13/cobra"
	"os"
)

var mainCmd = &cobra.Command{
	Use: "command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func main() {
	mainCmd.AddCommand(node.Cmd())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
