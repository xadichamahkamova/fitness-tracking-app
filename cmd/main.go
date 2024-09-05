package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
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

	ctx := context.Background()
	queries := storage.New(db)
	
	user, err := queries.CreateUser(ctx, storage.CreateUserParams{
		Username:     sql.NullString{String: "test", Valid: true},
		Email:        sql.NullString{String: "test@gmail.com", Valid: true},
		Profile:      pqtype.NullRawMessage{},
		PasswordHash: sql.NullString{String: "hashedpassword", Valid: true},
	})
	if err != nil {
		fmt.Println("failed to create user:", err)
		os.Exit(1)
	}
	fmt.Println("User created with id: ", user.ID)

	users, err := queries.ListUsers(ctx)
	if err != nil {
		logger.Error("failed to create user")
		os.Exit(1)
	}
	for _, v := range users {
		s := v.Profile.RawMessage
		fmt.Printf("user: %+v\n", string(s))
	}

	userByID, err := queries.GetUser(ctx, user.ID)
	if err != nil {
		fmt.Println("failed to get user by ID:", err)
		os.Exit(1)
	}
	fmt.Printf("Fetched user: %+v\n", userByID)

	err = queries.UpdateUser(ctx, storage.UpdateUserParams{
		ID:       user.ID,
		Username: sql.NullString{String: "updated_username", Valid: true},
		Email:    sql.NullString{String: "updated_email@gmail.com", Valid: true},
		Profile:  pqtype.NullRawMessage{},
	})
	if err != nil {
		fmt.Println("failed to update user:", err)
		os.Exit(1)
	}
	fmt.Println("User updated successfully")


}