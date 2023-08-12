package configure

import (
	"fmt"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Display the organization currently set as active",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement the get logic here
		fmt.Println("Getting a property...")
	},
}
