package configure

import (
	"fmt"
	"github.com/spf13/cobra"
	"warestack-cli-v2/pkg/ui"
	"warestack-cli-v2/pkg/util"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Display the organization currently set as active",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := util.ConfigFilePath(configFile)
		if err != nil {
			fmt.Println("Error:", err)
		}
		var config Config
		err = util.ReadJSON(path, &config)
		if err != nil {
			fmt.Println("Error:", err)
		}

		ui.ShowFormattedMessage("Your current organization alias is [%s]. You can switch between organizations by running:\n", config.CurrentOrganization)
		ui.ShowMessage("$ warestack config set organization $ORGANIZATION_ALIAS")
		ui.ShowMessage("")
	},
}
