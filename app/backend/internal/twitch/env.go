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
	errEnvNotFound      string    = "Enviroment argument %s not found"
	enviromentArguments [6]string = [...]string{
		"TWITCH_CLIENT_ID",
		"TWITCH_CLIENT_SECRET",
		"USER_SERVICE_BASE_URL",
		"AUTH_SERVICE_BASE_URL",
		"SESSION_SERVICE_BASE_URL",
		"ACTION_SERVICE_BASE_URL",
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
