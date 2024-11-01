package sync

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
** This defines the expected enviroment variables for the Google Service
 */

var (
	errEnvNotFound      string     = "Enviroment argument %s not found"
	enviromentArguments [23]string = [...]string{
		"TOKEN_SECRET",
		"ADMIN_TOKEN",
		"USER_SERVICE_BASE_URL",
		"AUTH_SERVICE_BASE_URL",
		"AUTH_PORT",
		"SESSION_SERVICE_BASE_URL",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"TWITCH_CLIENT_ID",
		"TWITCH_CLIENT_SECRET",
		"GITHUB_KEY",
		"GITHUB_SECRET",
		"AUTH_KEY",
		"AUTH_MAX_AGES",
		"AUTH_IS_PROD",
		"WEB_BASE_URL",
		"WEB_PORT",
		"DISCORD_KEY",
		"DISCORD_SECRET",
		"SPOTIFY_KEY",
		"SPOTIFY_SECRET",
		"TWITCH_CLIENT_ID",
		"TWITCH_CLIENT_SECRET",
	}
)

func Env(envPath string) error {
	err := godotenv.Load(envPath)

	if err != nil {
		return err
	}
	for _, envArg := range enviromentArguments {
		v := os.Getenv(envArg)

		if v != "" {
			continue
		}
		return fmt.Errorf(errEnvNotFound, envArg)
	}
	return nil
}
