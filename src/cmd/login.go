package cmd

import "github.com/spf13/cobra"

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "OAuth2 Login",
	Run: func(cmd *cobra.Command, args []string) {
		handleLogin()
	},
}

func handleLogin() {
	print("hello")
}
