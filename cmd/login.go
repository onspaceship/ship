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

		client := client.NewClient()

		_, err := client.GetUser()
		if err == nil {
			color.HiBlue("You're already logged in!")
			os.Exit(0)
		}

		url, err := client.GetTokenURL()
		if err != nil {
			log.Fatalf("Could not get a login URL: %v", err)
		}

		color.White("Opening up %s in your browser...", url)
		browser.OpenURL(url)

		urlParts := strings.Split(url, "/")
		code := urlParts[len(urlParts)-1]
		tokenStr, err := fetchToken(ctx, client, code)
		if err != nil {
			color.HiRed(err.Error())
		} else {
			token.SaveToken(tokenStr)
			color.HiBlue("You are now logged into Spaceship! 🚀")
		}
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
