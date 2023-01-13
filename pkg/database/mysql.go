package database

import (
	"database/sql"
	"fmt"
	"github.com/fahturr/xsis-test/config"
	"os"
	"time"
)

func VerifyDBConnection() {
	isConnected := false
	for !isConnected {
		_, err := ConnectSQLDB()
		if err == nil {
			isConnected = true
		}

		time.Sleep(2500 * time.Millisecond)
	}
}

func ConnectSQLDB() (*sql.DB, error) {
	var (
		DB_USER = os.Getenv(config.SQL_USER)
		DB_PASS = os.Getenv(config.SQL_PASS)
		DB_HOST = os.Getenv(config.SQL_HOST)
		DB_PORT = os.Getenv(config.SQL_PORT)
		DB_NAME = os.Getenv(config.SQL_DB)
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
