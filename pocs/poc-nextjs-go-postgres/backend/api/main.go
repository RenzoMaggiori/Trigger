package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"trigger.com/api/src/database"
	"trigger.com/api/src/parser"
	"trigger.com/api/src/server"
)

func main() {
	args, err := parser.CmdArgs()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("args", args)

	err = godotenv.Load(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Environ())

	database, err := database.Open(database.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	server := server.Create(*args.Port)
	go server.Start()
	defer server.Stop()
}
