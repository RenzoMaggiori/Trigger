package spotify

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
** This defines the expected enviroment variables for the Spotify Service
 */

var (
	errEnvNotFound      string    = "Enviroment argument %s not found"
	enviromentArguments [6]string = [...]string{
		"SESSION_SERVICE_BASE_URL",
		"SYNC_SERVICE_BASE_URL",
		"USER_SERVICE_BASE_URL",
		"ACTION_SERVICE_BASE_URL",
		"WORKSPACE_SERVICE_BASE_URL",
		"ADMIN_TOKEN",
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
