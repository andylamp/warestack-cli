package login

import (
	"fmt"
	"github.com/spf13/cobra"
	"warestack-cli-v2/pkg/api"
	"warestack-cli-v2/pkg/auth"
	"warestack-cli-v2/pkg/ui"
)

// AuthURL is the URL the user will be sent to for authentication
const AuthURL = "https://www.warestack.cloud/login"

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
		auth.OpenBrowser(AuthURL)
		<-done
		ShutDownServer(server)

		tokenClaims, err := auth.ReadTokenClaims()
		client := api.NewClient(tokenClaims)
		user, err := client.GetUserProfile()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		organization, err := client.GetOrganizationByID(user.Settings.DefaultOrganizationID)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		ui.ShowFormattedMessage("You are now logged in as [%s]!\n", user.Email)
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
