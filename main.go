package main

import (
	"github.com/oexlkinq/fairdrop/internal/config"
	"github.com/oexlkinq/fairdrop/internal/db"
	"github.com/oexlkinq/fairdrop/internal/routes"
	"github.com/oexlkinq/fairdrop/internal/storage"
)

func main() {
	config := config.Populate()

	db := db.Connect(config.DataFolderPath)
	s := storage.Create(config.DataFolderPath)

	routes := routes.Create(db, s)
	routes.Run()
}
