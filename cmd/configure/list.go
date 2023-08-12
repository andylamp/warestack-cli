package configure

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the organizations you're a part of",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement the get logic here
		fmt.Println("Getting a property...")
	},
}
