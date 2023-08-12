package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"warestack-cli-v2/cmd/configure"
	"warestack-cli-v2/cmd/login"
)

var rootCmd = &cobra.Command{
	Use:   "war",
	Short: "Warestack CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ERROR: Command name argument expected.")
		cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(login.Cmd, configure.Cmd)
	rootCmd.Execute()
}
