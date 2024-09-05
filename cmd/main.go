package main

import (
	"database/sql"
	"log/slog"
	"os"
	_ "github.com/lib/pq"
)

func main() {

	connString := "postgres://postgres:1@localhost:5432/postgres?sslmode=disable"
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Error("failed to connect")
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("failed to ping")
		os.Exit(1)
	}
}