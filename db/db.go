package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var Bun *bun.DB

func CreateDatabase(
	dbname string,
	user string,
	dbpassword string,
	dbhost string,
	dbport string,
) (*sql.DB, error) {
	hostArr := strings.Split(dbhost, ":")
	host := hostArr[0]
	port := dbport
	if len(hostArr) > 1 {
		port = hostArr[1]
	}
	uri := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user,
		dbpassword,
		dbname,
		host,
		port,
	)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Init() error {
	var (
		host   = os.Getenv("DB_HOST")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
		port   = os.Getenv("DB_PORT")
	)
	db, err := CreateDatabase(dbname, user, pass, host, port)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	Bun = bun.NewDB(db, pgdialect.New())

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return nil
}
