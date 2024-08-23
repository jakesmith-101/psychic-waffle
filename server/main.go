package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

// urlExample := "postgres://username:password@localhost:5432/database_name"
func buildDBUrl(dbType string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s",
		os.Getenv(fmt.Sprintf("%sUSER", dbType)),
		os.Getenv(fmt.Sprintf("%sPASSWORD", dbType)),
		os.Getenv(fmt.Sprintf("%sPORT", dbType)),
		os.Getenv(fmt.Sprintf("%sNAME", dbType)),
	)
}

func main() {
	conn, err := pgx.Connect(context.Background(), buildDBUrl("TEST_DB_")) // prod: "DB_", test: "TEST_DB_"
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}
