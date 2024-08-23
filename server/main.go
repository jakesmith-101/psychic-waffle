package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jakesmith-101/psychic-waffle/db"
)

func main() {
	db.Open()
	defer db.Conn.Close(context.Background())

	router := mux.NewRouter()
	router.HandleFunc("/users", db.GetUsers).Methods("GET")
	router.HandleFunc("/users", db.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
