package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CtxKey string = "ctxDatabaseKey"

func Open(connectionString string) (*mongo.Client, context.Context, error) {
	ctx := context.TODO()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		return nil, nil, err
	}
	if err := db.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}
	return db, ctx, nil
}

func ConnectionString() string {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	port := os.Getenv("MONGO_PORT")
	host := os.Getenv("MONGO_HOST")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", user, pass, host, port)
}
