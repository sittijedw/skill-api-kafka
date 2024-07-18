package database

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Println("Error: Can't connect to database", err)
	} else {
		log.Println("Connect database success")
	}

	return db
}
