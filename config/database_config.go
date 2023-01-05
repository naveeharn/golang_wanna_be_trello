package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic("Fasiiled to load .env file")
	}

	pg_username := os.Getenv("POSTGRES_USERNAME")
	pg_password := os.Getenv("POSTGRES_PASSWORD")
	pg_hostname := os.Getenv("POSTGRES_HOSTNAME")
	pg_port := os.Getenv("POSTGRES_PORT")
	pg_db_name := os.Getenv("POSTGRES_DB_NAME")

	postgres_url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		pg_username, pg_password, pg_hostname, pg_port, pg_db_name,
	)

	db, err := gorm.Open(postgres.Open(postgres_url), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to postgres database")
	}

	err = db.AutoMigrate()

	if err != nil {
		panic("Failed to migrate from entities to postgres database")
	}
	log.Println("Succeeded to connect postgres database")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPostgres, err := db.DB()
	if err != nil {
		panic("Failed to close connection from postgres database")
	}
	dbPostgres.Close()
}
