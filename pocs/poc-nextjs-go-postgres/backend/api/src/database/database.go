package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Open(connectionString string) (*sql.DB, error) {
	database, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	if err := database.Ping(); err != nil {
		return nil, err
	}
	return database, nil
}

func ConnectionString() string {
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")
	postgres_host := os.Getenv("POSTGRES_HOST")
	postgres_port := os.Getenv("POSTGRES_PORT")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", postgres_user, postgres_password, postgres_host, postgres_port, postgres_db)
}
