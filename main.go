package main

import (
	"context"
	"log"

	"github.com/Samuelmasih6/Bankcore/api"
	"github.com/Samuelmasih6/Bankcore/util"

	db "github.com/Samuelmasih6/Bankcore/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("No .env file found")
	// }

	// dbSource := os.Getenv("DB_SOURCE")
	// if dbSource == "" {
	// 	log.Fatal("DB_SOURCE is not set")
	// }
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect to PostgreSQL
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer conn.Close()

	// Create store
	store := db.NewStore(conn)

	// Create API server
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	// Start server
	log.Printf("starting server at %s\n", config.ServerAddress)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
