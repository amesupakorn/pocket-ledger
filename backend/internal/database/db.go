package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
