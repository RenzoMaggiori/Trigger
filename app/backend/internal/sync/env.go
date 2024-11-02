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
	enviromentArguments [26]string = [...]string{
		"TOKEN_SECRET",
		"ADMIN_TOKEN",
		"AUTH_PORT",
		"AUTH_KEY",
		"AUTH_IS_PROD",
		"AUTH_MAX_AGES",
		"USER_SERVICE_BASE_URL",
		"AUTH_SERVICE_BASE_URL",
		"SESSION_SERVICE_BASE_URL",
		"WEB_BASE_URL",
		"WEB_PORT",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"TWITCH_CLIENT_ID",
		"TWITCH_CLIENT_SECRET",
		"GITHUB_KEY",
		"GITHUB_SECRET",
		"DISCORD_KEY",
		"DISCORD_SECRET",
		"BOT_TOKEN",
		"SPOTIFY_KEY",
		"SPOTIFY_SECRET",
		"SETTINGS_SERVICE_BASE_URL",
		"TWITCH_CLIENT_ID",
		"TWITCH_CLIENT_SECRET",
		"SERVER_BASE_URL",
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
