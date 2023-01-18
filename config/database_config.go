package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	// log.Println(filepath.Join(".env"))
	// path, err := os.Getwd()
	// helper.LoggerErrorPath(runtime.Caller(0))
	// if err != nil {
	// log.Println(err)
	// }
	// log.Println(path)

	// fname := "../main.go"
	// abs_fname, err := filepath.Abs(fname)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(abs_fname)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// file executable
	// fmt.Println(ex)

	// the executable directory
	exPath := filepath.Dir(ex)
	// fmt.Println(exPath)
	// b, err := ioutil.ReadFile(os.Args[1])
	log.Println(exPath + "/.env")
	log.Println(filepath.Join(".env"))

	if err := godotenv.Load(exPath + "/.env"); err != nil {
		panic("Failed to load .env file")
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

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Team{},
		&entity.Dashboard{},
		&entity.Note{},
	)

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
