package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"trigger.com/api/src/database"
	"trigger.com/api/src/parser"
	"trigger.com/api/src/server"
)

func main() {
	args, err := parser.CmdArgs()

	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	database, err := database.Open(database.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	server, err := server.Create(*args.Port, database)
	if err != nil {
		log.Fatal(err)
	}

	go server.Start()
	defer server.Stop()
}
