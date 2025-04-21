package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataFolderPath string
	BasePath       string
}

func Populate() Config {
	godotenv.Load()

	// #3 TODO: чекать, существует ли, хотябы пустая, переменная в энве
	return Config{
		DataFolderPath: os.Getenv("DATA_FOLDER_PATH"),
		BasePath:       os.Getenv("BASE_PATH"),
	}
}
