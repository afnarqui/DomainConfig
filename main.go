package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

var Host string
var db *sql.DB

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8001"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "vertion initial: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":"+PORT, nil)
}

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}

