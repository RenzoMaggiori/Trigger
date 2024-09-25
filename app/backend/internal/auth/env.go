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
	enviromentArguments [14]string = [...]string{
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"GITHUB_KEY",
		"GITHUB_SECRET",
		"AUTH_KEY",
		"AUTH_MAX_AGES",
		"AUTH_IS_PROD",
		"MONGO_INITDB_ROOT_USERNAME",
		"MONGO_INITDB_ROOT_PASSWORD",
		"MONGO_PORT",
		"MONGO_HOST",
		"MONGO_DB",
		"TOKEN_SECRET",
		"USER_SERVICE_BASE_URL",
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
