package database

import (
	"database/sql"
	"log"
)

func Conn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Hacker"
	dbName := "parking"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Panicln("Error in DB Connection ", err.Error())
	}
	return db
}
