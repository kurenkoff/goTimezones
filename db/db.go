package db

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Database struct {
	db *sql.DB
}

func New() *Database {
	_ = pq.Efatal
	connStr := "user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD")
	connStr += " dbname=" + os.Getenv("POSTGRES_DB")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{
		db,
	}
}
