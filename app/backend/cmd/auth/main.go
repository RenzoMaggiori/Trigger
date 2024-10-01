package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"trigger.com/trigger/internal/auth"
	"trigger.com/trigger/internal/auth/credentials"
	"trigger.com/trigger/internal/auth/providers"
	"trigger.com/trigger/pkg/arguments"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/router"
	"trigger.com/trigger/pkg/server"
)

func main() {
	args, err := arguments.Command()
	if err != nil {
		log.Fatal(err)
	}

	err = auth.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	callback := fmt.Sprintf("%s/api/oauth2/callback", os.Getenv("AUTH_SERVICE_BASE_URL"))
	providers.CreateProvider(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			callback,
		),
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			callback,
		),
	)

	router, err := router.Create(
		context.TODO(),
		credentials.Router,
		providers.Router,
	)
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.Create(
		router,
		middleware.Create(
			middleware.Logging,
			middleware.Cors,
		),
		*args.Port,
	)
	if err != nil {
		log.Fatal(err)
	}

	go server.Start()
	server.Stop()
}
