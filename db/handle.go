package db

import (
	"database/sql"
	"fmt"
	"mail-sender/tools"
	"time"
)

var Handle *sql.DB

func Connect() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		tools.GetEnv("DB_USERNAME", ""),
		tools.GetEnv("DB_PASSWORD", ""),
		tools.GetEnv("DB_HOSTNAME", ""),
		tools.GetEnv("DB_PORT", "3306"),
		tools.GetEnv("DB_SCHEMA", "newtendex"),
	))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Hour * 4)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)

	Handle = db

	return
}
