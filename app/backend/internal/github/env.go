package github

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
** This defines the expected enviroment variables for the Auth Service
 */

var (
	errEnvNotFound      string    = "Enviroment argument %s not found"
	enviromentArguments [8]string = [...]string{
		"MONGO_INITDB_ROOT_USERNAME",
		"MONGO_INITDB_ROOT_PASSWORD",
		"MONGO_PORT",
		"MONGO_HOST",
		"MONGO_DB",
		"GITHUB_CLIENT_ID",
		"GITHUB_CLIENT_SECRET",
		"GITHUB_SERVICE_BASE_URL",
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
