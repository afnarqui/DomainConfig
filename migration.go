package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB
func main() {

	fmt.Println("running the migration")

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	if _, err := db.Exec(
		`DROP TABLE IF EXISTS domain`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`DROP TABLE IF EXISTS domainold`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`DROP TABLE IF EXISTS domainhistory`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS domain (
					Host VARCHAR(120) NULL,
					Port INT NULL,
					Protocol VARCHAR(120) NULL,
					IsPublic BOOL NULL,
					Status   VARCHAR(80) NULL,
					Endpoints       VARCHAR(8000) NULL 
					)`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS domainold (
				Host VARCHAR(120) NULL,
				Port INT NULL,
				Protocol VARCHAR(120) NULL,
				IsPublic BOOL NULL,
				Status   VARCHAR(80) NULL,
				Endpoints       VARCHAR(8000) NULL 
				)`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS domainhistory (
				Host VARCHAR(120) NULL,
				Port INT NULL,
				Protocol VARCHAR(120) NULL,
				IsPublic BOOL NULL,
				Status   VARCHAR(80) NULL,
				Endpoints       VARCHAR(8000) NULL 
				)`); err != nil {
		log.Fatal(err)
	}
}
