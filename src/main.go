package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/go-auth-cli/src/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "authcli",
	Short: "Go cli with OAuth2 login",
}

func init() {
	rootCmd.AddCommand(cmd.LoginCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
