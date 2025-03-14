package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBFile         string
	PublicBasePath string
	StoragePath    string
}

func Populate() Config {
	godotenv.Load()

	// #3 TODO: чекать, существует ли, хотябы пустая, переменная в энве
	return Config{
		DBFile:         os.Getenv("DB_FILE"),
		PublicBasePath: os.Getenv("PUBLIC_BASE_PATH"),
		StoragePath:    os.Getenv("STORAGE_PATH"),
	}
}
