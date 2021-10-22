package cmd

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/onspaceship/ship/pkg/client"
	"github.com/onspaceship/ship/pkg/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultRegistryHost = "registry.onspaceship.com"
)

var configureDockerCmd = &cobra.Command{
	Use:   "configure-docker",
	Short: "Sets up Docker authentication to the Magic Container Registry",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.CommandExists("docker") {
			color.HiRed("Unable to find docker executable.")
			os.Exit(1)
		}

		dockerCommand := exec.Command("docker", "login", "--username", "token", "--password-stdin", viper.GetString("registry_host"))
		dockerCommand.Stdout = os.Stdout
		dockerCommand.Stderr = os.Stderr

		stdin, err := dockerCommand.StdinPipe()
		if err != nil {
			log.Fatalf("Problem getting stdin to docker command: %v", err)
		}

		client := client.NewClient()
		token, err := client.GetRegistryToken(viper.GetString("current_team"))
		if token == "" || err != nil {
			color.HiRed("Unable to get authentication token. Are you logged in?")
			os.Exit(1)
		}

		if err = dockerCommand.Start(); err != nil {
			log.Fatalf("Problem starting docker command: %v", err)
		}

		io.WriteString(stdin, token+"\n")
		stdin.Close()
		dockerCommand.Wait()

		if dockerCommand.ProcessState.ExitCode() == 0 {
			color.HiBlue("Docker is authenticated and ready to pull images from the Magic Container Registry! 🚀")
		}
	},
}

func init() {
	viper.SetDefault("registry_host", defaultRegistryHost)

	rootCmd.AddCommand(configureDockerCmd)
}
