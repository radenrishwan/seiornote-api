package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"seiornote/helper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := ""

	if os.Getenv("ENVIRONMENT") == "production" {
		database = os.Getenv("DB_DATABASE")
	} else {
		database = os.Getenv("DB_DATABASE_TEST")
	}

	// postgres dsn
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, host, port, database)

	// mysql dsn
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("postgres", dsn)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	return db
}

func Migration() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := ""

	if os.Getenv("ENVIRONMENT") == "production" {
		database = os.Getenv("DB_DATABASE")
	} else {
		database = os.Getenv("DB_DATABASE_TEST")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	m, err := migrate.New("file://database/scheme", dsn)

	if err != nil {
		log.Fatalln(err)
	}

	// migrate table
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			log.Fatalln(err.Error())
		}
	}

	log.Println("Migration success")

}
