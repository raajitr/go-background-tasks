package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBClient() *pgxpool.Pool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgxpool.New(context.Background(), "postgres://admin:admin@postgres:5432/mydb?sslmode=disable")
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}