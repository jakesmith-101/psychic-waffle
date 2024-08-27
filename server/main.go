package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" // replace with go-chi
	// add go-swagger for restful-ness

	"github.com/jakesmith-101/psychic-waffle/db"
)

func main() {
	db.Open()

	router := mux.NewRouter()
	router.HandleFunc("/users", db.GetUsers).Methods("GET")
	router.HandleFunc("/users", db.CreateUser).Methods("POST")

	fmt.Fprintf(os.Stderr, "Listening on port: 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
