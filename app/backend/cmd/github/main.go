package main

import (
	"context"
	"log"

	"trigger.com/trigger/internal/github"
	"trigger.com/trigger/internal/github/reaction"
	"trigger.com/trigger/internal/github/trigger"
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

	err = github.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	router, err := router.Create(
		context.TODO(),
		trigger.Router,
		reaction.Router,
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
