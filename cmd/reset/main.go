package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/oriiyx/arso/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	var (
		host   = os.Getenv("DB_HOST")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
		port   = os.Getenv("DB_PORT")
	)
	return db.CreateDatabase(dbname, user, pass, host, port)
}

func main() {
	database, err := createDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	tables := []string{
		"schema_migrations",
		"Message",
	}

	for _, table := range tables {
		query := fmt.Sprintf("drop table if exists %s ;", table)
		if _, err := database.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
