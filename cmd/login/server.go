package login

import (
	"context"
	"log"
	"net/http"
	"warestack-cli-v2/pkg/auth"
)

// LocalServerURL is the URL to be used for listening to the auth callback
const LocalServerURL = "localhost:8050"

// StartServer starts a local server and listens for a shutdown signal
func StartServer(done chan bool) *http.Server {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleRedirect(w, r)
		done <- true
	})

	server := &http.Server{Addr: LocalServerURL}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return server
}

// ShutDownServer shuts down the provided server
func ShutDownServer(server *http.Server) {
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Failed to shut down server: %v", err)
	}
}
