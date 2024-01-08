package database

import (
	"database/sql"
	"time"

	"github.com/irfanandriansyah1997/learn-api/internal/utils"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "my_user:my_password@tcp(localhost:3306)/learn_api_database")
	utils.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
