package configs

import (
	"os"

	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
)

type configs struct {
	PORT    string
	APP_ENV string
	DB_URI  string
	URL     string
}

var Config configs

func init() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: "CONFIG",
	})
	Config = configs{
		PORT:    os.Getenv("PORT"),
		APP_ENV: os.Getenv("APP_ENV"),
		DB_URI:  os.Getenv("DB_URI"),
		URL:     os.Getenv("URL"),
	}

	logger.Info("Loaded configs")
}
