package main

import (
	"context"
	"log"
	"os"

	githubAuth "github.com/markbates/goth/providers/github"
	googleAuth "github.com/markbates/goth/providers/google"

	"trigger.com/trigger/internal/providers/register"
	"trigger.com/trigger/internal/providers/register/google"
	"trigger.com/trigger/pkg/arguments"
	gothProviders "trigger.com/trigger/pkg/authenticator/providers"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/router"
	"trigger.com/trigger/pkg/server"
)

func main() {
	args, err := arguments.Command()
	if err != nil {
		log.Fatal(err)
	}

	err = register.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	gothProviders.CreateProvider(
		googleAuth.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:8080/api/google/register/callback"),
		githubAuth.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			"http://localhost:8080/api/github/register/callback"),
	)

	router, err := router.Create(
		context.TODO(),
		google.Router,
		register.Router,
		//github.Router,
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
