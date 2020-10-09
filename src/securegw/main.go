package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	port = "9943"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ping", ping).Methods("GET")
	log.Println("Starting http server on port: ", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
