package cmd

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/onspaceship/ship/pkg/client"
	"github.com/onspaceship/ship/pkg/token"

	"github.com/fatih/color"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	tokenRequestTimeout  = time.Minute * 2
	tokenRequestDuration = time.Second * 1
)

var loginCmd = &cobra.Command{
	Use:                   "login",
	Short:                 "Log into your Spaceship account",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), tokenRequestTimeout)
		defer cancel()

		core := client.NewClient()

		_, err := core.GetUser()
		if err == nil {
			color.HiBlue("You're already logged in!")
			os.Exit(0)
		}

		url, err := core.GetTokenURL()
		if err != nil {
			log.Fatalf("Could not get a login URL: %v", err)
		}

		color.White("Opening up %s in your browser...", url)
		browser.OpenURL(url)

		urlParts := strings.Split(url, "/")
		code := urlParts[len(urlParts)-1]
		tokenStr, err := fetchToken(ctx, core, code)
		if err != nil {
			log.Fatalf("Could not get a token: %v", err)
		}

		token.SaveToken(tokenStr)
		core = client.NewClient()

		user, err := core.GetUser()
		if err != nil {
			log.Fatal(err)
		}

		if len(user.Teams) == 0 {
			color.HiRed("You must be a member of at least one team.")
			token.ClearToken()
			os.Exit(1)
		}

		viper.Set("current_team", user.Teams[0].Handle)
		if err = viper.WriteConfig(); err != nil {
			log.Fatal(err)
		}

		color.HiBlue("You are now logged into Spaceship! 🚀")
	},
}

func fetchToken(ctx context.Context, client *client.Client, code string) (string, error) {
	ticker := time.NewTicker(time.Second * 2)

	for {
		select {
		case <-ticker.C:
			token, err := client.GetToken(code)
			if err == nil {
				return token, nil
			}
		case <-ctx.Done():
			return "", errors.New("timeout exceeded awaiting login from the browser")
		}
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
