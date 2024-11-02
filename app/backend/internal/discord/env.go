package discord

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
** This defines the expected enviroment variables for the Discord Service
 */

var (
	errEnvNotFound      string    = "Enviroment argument %s not found"
	enviromentArguments [14]string = [...]string{
		"MONGO_INITDB_ROOT_USERNAME",
		"MONGO_INITDB_ROOT_PASSWORD",
		"MONGO_PORT",
		"MONGO_HOST",
		"MONGO_DB",
		"SESSION_SERVICE_BASE_URL",
		"SYNC_SERVICE_BASE_URL",
		"USER_SERVICE_BASE_URL",
		"ACTION_SERVICE_BASE_URL",
		"ADMIN_TOKEN",
		"DISCORD_KEY",
		"DISCORD_SECRET",
		"BOT_TOKEN",
		"AUTH_SERVICE_BASE_URL",
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
