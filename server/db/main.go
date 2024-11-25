package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jakesmith-101/psychic-waffle/util"
)

var PgxPool *pgxpool.Pool
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
	PgxPool, err = pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		util.Log("error", "Unable to connect to database: %v", err)
		util.Log("error", "Database URL: %s", dbUrl)
		os.Exit(1)
	} else {
		util.Log("ouput", "Connected to database: %s", dbType)
	}

	err = DBTriggersFuncs()
	if err != nil {
		util.Log("error", "DB triggers and funcs creation failed: %v", err)
	} else {
		util.Log("ouput", "Created DB triggers and funcs")
	}
}

func DBTriggersFuncs() error {
	err := PostFuncs()
	return err
}
