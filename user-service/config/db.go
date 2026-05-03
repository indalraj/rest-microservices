package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:root@tcp(localhost:3306)/userdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
}

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(100)
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
