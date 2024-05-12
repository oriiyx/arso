package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/oriiyx/arso/db"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
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

	// Create migration instance
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Point to your migration files. Here we're using local files, but it could be other sources.
	migrationInstance, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // source URL
		"postgres",                      // database name
		driver,                          // database instance
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := migrationInstance.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := migrationInstance.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
