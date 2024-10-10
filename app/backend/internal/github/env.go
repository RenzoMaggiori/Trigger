package github

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
	enviromentArguments [1]string = [...]string{
		"WEBHOOK_URL",
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
