package login

import (
	"fmt"
	"github.com/spf13/cobra"
	"warestack-cli/cmd/configure"
	"warestack-cli/pkg/api"
	"warestack-cli/pkg/auth"
	"warestack-cli/pkg/ui"
	"warestack-cli/pkg/util"
)

// AuthURL is the URL the user will be sent to for authentication
const authURL = "https://www.warestack.com/cli-auth"
const credentialsFile = "credentials.json"
const configFile = "config.json"

var Cmd = &cobra.Command{
	Use:     "authenticate",
	Short:   "Authenticate the CLI via Warestack web application",
	Aliases: []string{"auth"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Channel to signal server shutdown
		done := make(chan bool)

		// Start the server and wait for the signal to shut it down
		server := StartServer(done)
		// open the web app login page in the default browser
		auth.OpenBrowser(authURL)
		<-done
		ShutDownServer(server)

		path, err := util.ConfigFilePath(credentialsFile)
		if err != nil {
			return err
		}
		var credentials auth.Credentials
		err = util.ReadJSON(path, &credentials)
		if err != nil {
			return err
		}

		client := api.NewClient(credentials)
		user, err := client.GetUserProfile()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		organization, err := client.GetOrganizationByAlias(user.Settings.DefaultOrganizationAlias)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		path, err = util.ConfigFilePath(configFile)
		if err != nil {
			return err
		}
		var config configure.Config
		config.CurrentOrganization = organization.Alias
		err = util.WriteJSON(path, &config)

		ui.ShowFormattedMessage("You are now authenticated as [%s]!\n", user.Email)
		ui.ShowMessage("")
		ui.ShowMessage("Success, The session will end after 7 days of inactivity. Otherwise, it will remain active.")
		ui.ShowMessage("")
		ui.ShowFormattedMessage("Your current organization alias is [%s]. You can switch between organizations by running:\n", organization.Alias)
		ui.ShowMessage("$ warestack config set organization $ORGANIZATION_ALIAS")
		ui.ShowMessage("")
		ui.ShowMessage("Enjoy Warestack!")
		return nil
	},
}
