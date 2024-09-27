package main

import (
	"context"
	"log"
	"os"

	googleAuth "github.com/markbates/goth/providers/google"
	"trigger.com/trigger/internal/google"
	"trigger.com/trigger/pkg/arguments"
	"trigger.com/trigger/pkg/authenticator/providers"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/router"
	"trigger.com/trigger/pkg/server"
)

func main() {
	args, err := arguments.Command()
	if err != nil {
		log.Fatal(err)
	}

	err = google.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	providers.CreateProvider(googleAuth.New(
		os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		"http://localhost:8000/api/auth/google/callback"))

	router, err := router.Create(
		context.WithValue(context.TODO(), providers.ProviderKey, "google/sync"),
		providers.ProviderRouter,
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
