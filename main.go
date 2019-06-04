package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/go-chi/chi"
	"time"
	"strings"
)

var Host string
var db *sql.DB

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	
	log.Println("Corriendo en http://localhost:8001")
	r := chi.NewRouter()

	fmt.Println(r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	direccion := ":8081" 
	fmt.Println("Servidor listo escuchando en " + direccion)

	log.Fatal(http.ListenAndServe(direccion+"/public/index.html", nil))
	
}

func Logger() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(time.Now(), r.Method, r.URL)
        router.ServeHTTP(w, r) 
    })
}
var router *chi.Mux

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
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

