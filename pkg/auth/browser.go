package auth

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"warestack-cli-v2/pkg/ui"
)

// OpenBrowser opens a browser window with the given URL.
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	ui.ShowMessage("Visit the following link to authorize the CLI\n")
	ui.ShowMessage(url)
	ui.ShowMessage("")

	if err != nil {
		log.Println("Failed to automatically open the URL. Please open it manually.")
	}
}
