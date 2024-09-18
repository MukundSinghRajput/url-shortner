package db

import (
	"database/sql"
	"os"

	"github.com/MukundSinghRajput/goblog/internals/configs"
	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

var (
	err error
	DB  *sql.DB
)

func NewDB() *sql.DB {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: "DATABASE",
	})
	DB, err = sql.Open("postgres", configs.Config.DB_URI)
	if err != nil {
		logger.Error("Error connecting to the database:", err)
		os.Exit(1)
	}
	if err := DB.Ping(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	DB.Exec(`
	CREATE TABLE IF NOT EXISTS tiny (
    origin text PRIMARY KEY,
    short text UNIQUE NOT NULL
);
`)
	logger.Info("Established connection with the database.")
	return DB
}
