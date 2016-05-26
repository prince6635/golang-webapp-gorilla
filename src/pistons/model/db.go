package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	database, err := sql.Open("postgres", "user=pistons_admin dbname=pistons password=admin sslmode=disable")
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
	}
}
