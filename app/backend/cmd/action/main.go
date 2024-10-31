package main

import (
	"context"
	"log"
	"os"

	"trigger.com/trigger/internal/action"
	actions "trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/worker"
	"trigger.com/trigger/internal/action/workspace"
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

	err = action.Env(*args.EnvPath)
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
	).Collection("workspace")

	ctx := context.WithValue(
		context.TODO(),
		actions.ActionCtxKey,
		actionCollection,
	)
	ctx = context.WithValue(
		ctx,
		workspace.WorkspaceCtxKey,
		userActionCollection,
	)

	router, err := router.Create(
		ctx,
		actions.Router,
		workspace.Router,
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
	if err := worker.Run(actionCollection); err != nil {
		log.Println(err)
	}
	server.Stop()
}
