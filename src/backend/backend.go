package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const port = ":8080"

func usernameHandler(w http.ResponseWriter, r *http.Request) {
	type User struct{ Username string }
	user := User{os.Getenv("USERNAME")}
	p, _ := json.Marshal(user)
	w.Write(p)
}

func main() {
	log.Println("Starting Backend")

	r := mux.NewRouter()
	// Define API routes
	r.HandleFunc("/api/username", usernameHandler).Methods("GET")

	// Serve webapp static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("dist")))

	log.Println("Http Listening")
	http.ListenAndServe(
		port, r)
}
