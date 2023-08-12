package configure

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [organization_alias]",
	Short: "Set a specific organization as the current active one",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement the get logic here
		fmt.Println("Getting a property...")
	},
}
