package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	// MsSql Import
	_ "github.com/denisenkom/go-mssqldb"
)

// Get DB Connection
func ConnectDB() (*sqlx.DB, error) {
	return GetTapConnection()
}

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

// Database Connector function
func GetTapConnection() (*sqlx.DB, error) {

	fmt.Println("[TAP-debug] Check Database State...")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Println(fmt.Sprintf("[TAP-debug] host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	tapDB, connError := sqlx.Open("postgres", psqlInfo)

	if connError != nil {
		log.Panicln("[TAP-debug] Error establishing connection to *TAP* database", connError)
		panic(connError)
	}

	pingError := tapDB.Ping()
	if pingError != nil {
		log.Panicln("[TAP-debug] Error connecting to *TAP* database", pingError)
		panic(pingError)
	}

	return tapDB, connError
}

func BeginTransact(db *sqlx.DB) (*sql.Tx, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	return tx, err
}
