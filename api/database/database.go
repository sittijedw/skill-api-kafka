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

func CreateSkillTable(db *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS skill (
			key TEXT PRIMARY KEY,
			name TEXT NOT NULL DEFAULT '',
			description TEXT NOT NULL DEFAULT '',
			logo TEXT NOT NULL DEFAULT '',
			TAGS TEXT [] NOT NULL DEFAULT '{}'
		)
	`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Panic("Error: Can't create skill table", err)
	} else {
		log.Println("Create skill table success")
	}
}
