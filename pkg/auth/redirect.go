package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HandleRedirect captures the Firebase tokens from the POST request body.
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var tokenClaims TokenClaims
	err := json.NewDecoder(r.Body).Decode(&tokenClaims)
	if err != nil {
		http.Error(w, "Failed to decode tokens", http.StatusBadRequest)
		return
	}

	err = storeTokenClaims(tokenClaims)
	if err != nil {
		log.Fatalf("Failed to write the tokens to file: %s", err)
		http.Error(w, "Failed to store tokens", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successful login. You can now return to the CLI!")
}
