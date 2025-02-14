package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// The postgreSQL connnection string
	databaseURL := "postgres://postgres:MR.ibrahim2001@localhost:5432/clinic"

	// Connect to postgresql
	conn, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		fmt.Println("failed to connect to Database", err)
	}
	defer conn.Close()

	// Test the connection
	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan((&version))
	if err != nil {
		fmt.Println("failed to fetch postgresql version", err)
	}
	fmt.Println("✅ Successfully connected to PostgreSQL!")
	fmt.Println("PostgreSQL Version:", version)

}
