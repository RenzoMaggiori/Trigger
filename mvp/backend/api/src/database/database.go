package database

import "log"

type Database struct{}

const CtxKey string = "ctxDatabaseKey"

// TODO: replace all the functions with the actual database
func Open(connectionString string) (*Database, error) {
	log.Println("Mock database connection opened")
	return &Database{}, nil
}

func (d *Database) Close() error {
	log.Println("Mock database connection closed")
	return nil
}

func ConnectionString() string {
	return "mock-db-string"
}
