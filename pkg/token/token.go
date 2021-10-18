package token

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
)

var (
	credStore    = viper.New()
	currentToken = ""
)

func GetToken() string {
	if currentToken != "" {
		return currentToken
	}

	token, err := keyring.Get("spaceship-cli", "apiToken")
	if err == nil {
		currentToken = token
		return token
	}

	currentToken = credStore.GetString("api_token")
	return currentToken
}

func SaveToken(token string) {
	err := keyring.Set("spaceship-cli", "apiToken", token)
	if err == nil {
		currentToken = token
		return
	}

	credStore.Set("api_token", token)
	err = credStore.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}

	currentToken = token
}

func HasToken() bool {
	isCleared := GetToken() != ""
	return isCleared
}

func ClearToken() {
	if HasToken() {
		SaveToken("")
	}
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := home + "/.ship/credentials.yaml"
	credStore.SetConfigFile(configPath)

	if err := credStore.SafeWriteConfigAs(configPath); err != nil {
		if os.IsNotExist(err) {
			if err = credStore.WriteConfigAs(configPath); err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := credStore.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}
}
