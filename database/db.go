package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connStr := "user=postgres password=postgre dbname=rpldb sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak merespons:", err)
	}

	log.Println("Terkoneksi ke PostgreSQL")
}
