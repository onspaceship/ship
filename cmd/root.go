package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ship",
	Short: "The command line tool for Spaceship 🚀",
	Long:  "The command line tool for Spaceship 🚀",
}

func Execute() {
	rootCmd.SetOut(color.Output)

	cobra.AddTemplateFunc("setBoldBlue", color.New(color.FgBlue, color.Bold).SprintFunc())
	cobra.AddTemplateFunc("setYellow", color.New(color.FgYellow).SprintFunc())
	cobra.AddTemplateFunc("setCyan", color.New(color.FgCyan).SprintFunc())

	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{setBoldBlue "USAGE"}}`,
		`{{.UseLine}}`, `{{.UseLine | setYellow}}`,
		`{{.CommandPath}} [command]{{end}}`, `{{.CommandPath | setYellow}} {{setYellow "[command]"}}{{end}}`,
		`Available Commands:`, `{{setBoldBlue "COMMANDS"}}`,
		`{{rpad .Name .NamePadding }} {{.Short}}`, `{{rpad .Name .NamePadding | setYellow}} {{.Short | setYellow}}`,
		`Aliases:`, `{{setBoldBlue "ALIASES"}}`,
		`{{.NameAndAliases}}`, `{{.NameAndAliases | setYellow}}`,
		`Flags:`, `{{setBoldBlue "FLAGS"}}`,
		`Global Flags:`, `{{setBoldBlue "GLOBAL FLAGS"}}`,
		`.FlagUsages | trimTrailingWhitespaces}}`, `.FlagUsages | trimTrailingWhitespaces | setYellow}}`,
		`{{.CommandPath}} [command] --help`, `{{.CommandPath | setCyan}} {{setCyan "[command] --help"}}`,
	).Replace(usageTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var (
	appHandle  string
	teamHandle string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&teamHandle, "team", "t", "", "Team handle")
	rootCmd.PersistentFlags().StringVarP(&appHandle, "app", "a", "", "App handle")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := home + "/.ship/config.yaml"
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}

	if err = viper.WriteConfig(); err != nil {
		log.Fatal(err)
	}

	viper.SetEnvPrefix("ship")
	viper.AutomaticEnv()

	viper.BindPFlag("current_team", rootCmd.PersistentFlags().Lookup("team"))
	viper.BindPFlag("current_app", rootCmd.PersistentFlags().Lookup("app"))
}
