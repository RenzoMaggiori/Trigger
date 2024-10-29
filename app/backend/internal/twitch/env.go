package twitch

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
** This defines the expected enviroment variables for the Github Service
 */

var (
	errEnvNotFound      string     = "Enviroment argument %s not found"
	enviromentArguments [12]string = [...]string{
		"TWITCH_CLIENT_ID",
		"TWITCH_CLIENT_SECRET",
		"TWITCH_BASE_URL",
		"USER_SERVICE_BASE_URL",
		"AUTH_SERVICE_BASE_URL",
		"SESSION_SERVICE_BASE_URL",
		"SYNC_SERVICE_BASE_URL",
		"ACTION_SERVICE_BASE_URL",
		"ADMIN_TOKEN",
		"AUTH_PORT",
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
