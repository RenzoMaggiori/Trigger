package main

import (
	"context"
	"log"
	"os"

	"trigger.com/trigger/internal/action/action"
	useraction "trigger.com/trigger/internal/action/user_action"
	"trigger.com/trigger/internal/user"
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

	err = user.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, _, err := mongodb.Open(mongodb.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	actionCollection := mongoClient.Database(
		os.Getenv("MONGO_DB"),
	).Collection("action")

	userActionCollection := mongoClient.Database(
		os.Getenv("MONGO_DB"),
	).Collection("user_action")

	ctx := context.WithValue(
		context.TODO(),
		action.ActionCtxKey,
		actionCollection,
	)
	ctx = context.WithValue(
		ctx,
		useraction.UserActionCtxKey,
		userActionCollection,
	)

	router, err := router.Create(
		ctx,
		action.Router,
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
