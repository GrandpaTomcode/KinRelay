package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	url := os.Getenv("KINRELAY_DB_URL")
	if url == "" {
		log.Fatal("KINRELAY_DB_URL not set")
	}

	var err error
	Pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("Unable to connect to the database: ", err)
	}

	log.Println("Connected to database")
}