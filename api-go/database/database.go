package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var Db orm.DB

func init() {
	dbHost, dbPort := os.Getenv("DB_HOST"), os.Getenv("DB_PORT")

	Db = pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Database: os.Getenv("DB_NAME"),
	})

	if Db == nil {
		log.Printf("failed to connect to Database")
		os.Exit(1)
	}
}
