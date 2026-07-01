package main

import (
	"context"
	"log"
	"os"

	"Bankcore/api"
	db "Bankcore/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

const serverAddress = "0.0.0.0:8080"

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE is not set")
	}

	// Connect to PostgreSQL
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer conn.Close()

	// Create store
	store := db.NewStore(conn)

	// Create API server
	server := api.NewServer(store)

	// Start server
	log.Printf("starting server at %s\n", serverAddress)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
