package configure

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "configure",
	Short:   "Manage and switch between your organizations",
	Aliases: []string{"config"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ERROR: (cli.configure) Command name argument expected.")
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(listCmd, setCmd, currentCmd)
}
