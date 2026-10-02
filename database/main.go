package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "221232"
	dbName   = "cine_track"
)

func Connect() (*pgx.Conn, error) {
	ctx := context.Background()

	// First connect to PostgreSQL's default "postgres" database.
	adminConn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/postgres",
			user,
			password,
			host,
			port,
		),
	)
	if err != nil {
		return nil, err
	}
	defer adminConn.Close(ctx)

	// Create our database if it doesn't exist.
	var exists bool

	err = adminConn.QueryRow(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)",
		dbName,
	).Scan(&exists)

	if err != nil {
		return nil, err
	}

	if !exists {
		_, err = adminConn.Exec(
			ctx,
			fmt.Sprintf("CREATE DATABASE %s", dbName),
		)
		if err != nil {
			return nil, err
		}

		fmt.Println("Database created!")
	}

	// Now connect to our actual database.
	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s",
			user,
			password,
			host,
			port,
			dbName,
		),
	)
	if err != nil {
		return nil, err
	}

	// Create users table if it doesn't exist.
	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		)
	`)

	if err != nil {
		conn.Close(ctx)
		return nil, err
	}

	fmt.Println("Users table ready!")

	return conn, nil
}
