package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbPort == "" {
		dbPort = "5432"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic("could not ping database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createPropertiesTable := `
	CREATE TABLE IF NOT EXISTS properties (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		location TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createPropertiesTable)
	if err != nil {
		panic("could not create properties table: " + err.Error())
	}
}
