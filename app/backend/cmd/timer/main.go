package main

import (
	"context"
	"log"

	"trigger.com/trigger/internal/timer"
	"trigger.com/trigger/internal/timer/trigger"
	"trigger.com/trigger/internal/timer/worker"
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

	err = timer.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.TODO())

	router, err := router.Create(
		ctx,
		trigger.Router,
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
	worker := worker.Create(ctx, cancel)
	go worker.Start()

	server.Stop()
	worker.Stop()
}
