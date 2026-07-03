package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Samuelmasih6/Bankcore/util"
	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/joho/godotenv"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Println("cannot load config: ", err)
	}

	//dbSource := os.Getenv("DB_SOURCE")

	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()

	os.Exit(code)
}
