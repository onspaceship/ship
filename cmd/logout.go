package cmd

import (
	"github.com/onspaceship/ship/pkg/token"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:                   "logout",
	Short:                 "Log out of your Spaceship account",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		token.ClearToken()
		color.HiBlue("You are now logged out of Spaceship! 🚀")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
