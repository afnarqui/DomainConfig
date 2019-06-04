package main
import (
 "fmt"
 "net/http"
 "os"
)
func main() {
 var PORT string
 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "8001"
 }
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "vertion initial: %s\n", r.URL.Path)
 })
 http.ListenAndServe(":" + PORT, nil)
}