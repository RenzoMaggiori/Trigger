package auth

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
	enviromentArguments [12]string = [...]string{
		"TOKEN_SECRET",
		"ADMIN_TOKEN",
		"USER_SERVICE_BASE_URL",
		"AUTH_SERVICE_BASE_URL",
		"SESSION_SERVICE_BASE_URL",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"GITHUB_KEY",
		"GITHUB_SECRET",
		"AUTH_KEY",
		"AUTH_MAX_AGES",
		"AUTH_IS_PROD",

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
