package main

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully :)")
}
