package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"warestack-cli/pkg/util"
)

const credentialsFile = "credentials.json"

// HandleRedirect captures the Firebase tokens from the POST request body.
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Failed to decode tokens", http.StatusBadRequest)
		return
	}

	path, err := util.ConfigFilePath(credentialsFile)
	if err != nil {
		return
	}
	err = util.WriteJSON(path, credentials)
	if err != nil {
		log.Fatalf("Failed to write the tokens to file: %s", err)
		http.Error(w, "Failed to store tokens", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successful sign in. You can now return to the CLI!")
}
