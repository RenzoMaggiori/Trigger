package main

import (
	"context"
	"log"
	"os"

	"trigger.com/trigger/internal/spotify"
	"trigger.com/trigger/internal/spotify/reaction"
	"trigger.com/trigger/internal/spotify/trigger"
	"trigger.com/trigger/internal/spotify/worker"
	"trigger.com/trigger/pkg/arguments"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
	"trigger.com/trigger/pkg/server"
)

func main() {
	args, err := arguments.Command()
	if err != nil {
		log.Fatal(err)
	}

	err = spotify.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, _, err := mongodb.Open(mongodb.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	spotifyCollection := mongoClient.Database(
		os.Getenv("MONGO_DB"),
	).Collection("spotify")

	ctx := context.WithValue(
		context.TODO(),
		mongodb.CtxKey,
		spotifyCollection,
	)

	router, err := router.Create(
		ctx,
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
	worker := worker.New(ctx)
	worker.Start()

	server.Stop()
	worker.Stop()
}
