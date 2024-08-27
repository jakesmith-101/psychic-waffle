package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5" // keep or transition to gorm?
)

var Conn *pgx.Conn
var err error

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
func buildDBUrl(dbType string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv(fmt.Sprintf("%sUSER", dbType)),
		os.Getenv(fmt.Sprintf("%sPASSWORD", dbType)),
		os.Getenv(fmt.Sprintf("%sHOST", dbType)),
		os.Getenv(fmt.Sprintf("%sPORT", dbType)),
		os.Getenv(fmt.Sprintf("%sDB", dbType)),
	)
}

func Open() {
	dbType := "TEST_POSTGRES"                       // prod: "POSTGRES", test: "TEST_POSTGRES"
	dbUrl := buildDBUrl(fmt.Sprintf("%s_", dbType)) // apply connecting "_"
	Conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		fmt.Fprintf(os.Stderr, "Database URL: %s\n", dbUrl)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "Connected to database: %s\n", dbType)
		defer Conn.Close(context.Background())
	}
}
