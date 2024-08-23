package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/jakesmith-101/psychic-waffle/db"
)

func main() {
	time.Sleep(30 * time.Second)
	db.Open()
	defer db.Conn.Close(context.Background())

	router := mux.NewRouter()
	router.HandleFunc("/users", db.GetUsers).Methods("GET")
	router.HandleFunc("/users", db.CreateUser).Methods("POST")

	fmt.Fprintf(os.Stderr, "Listening on port: 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
