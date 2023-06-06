package utils

import (
	"database/sql"
	"fmt"
	"os"
)

func DatabaseConnector() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error when calling database!", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error when ping database!", err)
		return nil, err
	}
	return db, nil
}
