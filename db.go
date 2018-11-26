package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectDatabase() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "banking"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	fmt.Println("Database connected.")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
