package user

import (
	"fmt"
	"os"
)

/*
** This package defind the expected enviroment variables for the User Service
 */

var (
	errEnvNotFound      string    = "Enviroment argument %s not found"
	enviromentArguments [4]string = [...]string{
		"MONGO_INITDB_ROOT_USERNAME",
		"MONGO_INITDB_ROOT_PASSWORD",
		"MONGO_PORT",
		"MONGO_HOST",
	}
)

func Env() error {
	for _, envArg := range enviromentArguments {
		v := os.Getenv(envArg)

		if v != "" {
			continue
		}
		return fmt.Errorf(errEnvNotFound, v)
	}

	return nil
}
