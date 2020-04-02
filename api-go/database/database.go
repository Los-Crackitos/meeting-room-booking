package database

import (
	"log"
	"os"
	"fmt"
	
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/joho/godotenv"
)

var Db orm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}
	
	dbHost, dbPort := os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_PORT")


	Db = pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_DB_USER"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Database: os.Getenv("POSTGRES_DB_NAME"),
	})

	if Db == nil {
		log.Printf("failed to connect to Database")
		os.Exit(1)
	}
}
