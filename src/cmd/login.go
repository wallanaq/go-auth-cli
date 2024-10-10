package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/wallanaq/go-auth-cli/src/browser"
	"github.com/wallanaq/go-auth-cli/src/server"
	"golang.org/x/oauth2"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "OAuth2 Login",
	Run: func(cmd *cobra.Command, args []string) {
		handleLogin()
	},
}

func handleLogin() {

	oauth2Config := &oauth2.Config{
		ClientID:     "auth-cli",
		ClientSecret: "Er7HhGr9cIDHgYX4gmGHJQ7NMxLZEZLp",
		RedirectURL:  "http://localhost:8088/callback",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:8085/realms/dev/protocol/openid-connect/auth",
			TokenURL: "http://localhost:8085/realms/dev/protocol/openid-connect/token",
		},
	}

	state := "random-state"

	authURL := oauth2Config.AuthCodeURL(state)

	if err := browser.Open(authURL); err != nil {
		log.Fatalf("Failed to open browser: %v", err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")

		if code == "" {
			http.Error(w, "Missing code", http.StatusBadRequest)
			return
		}

		token, err := oauth2Config.Exchange(context.Background(), code)

		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Login successful! You can close your browser now.")

		log.Println("Login successful!")
		log.Println(token.AccessToken)

		server.NotifyDone()

	})

	server.Start(":8088", nil)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-server.Done():
	case <-sigs:
	}

	server.Shutdown()

	os.Exit(0)

}
