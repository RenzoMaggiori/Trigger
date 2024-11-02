package main

import (
	"context"
	"log"
	"os"

	"trigger.com/trigger/internal/discord"
	"trigger.com/trigger/internal/discord/reaction"
	"trigger.com/trigger/internal/discord/trigger"
	"trigger.com/trigger/internal/discord/worker"
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

	err = discord.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, _, err := mongodb.Open(mongodb.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	discordCollection := mongoClient.Database(
		os.Getenv("MONGO_DB"),
	).Collection("discord")

	ctx := context.WithValue(
		context.TODO(),
		mongodb.CtxKey,
		discordCollection,
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

	w := &worker.Model{}
	w.InitBot()
	// worker.Start()

	server.Stop()
	// worker.Stop()
}
