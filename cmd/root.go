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

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(home + "/.ship")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("ship")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}
}
