package main

import "os"

type Config struct {
	DBFile         string
	PublicBasePath string
}

func Populate() Config {
	return Config{
		DBFile:         os.Getenv("DB_FILE"),
		PublicBasePath: os.Getenv("PUBLIC_BASE_PATH"),
	}
}
