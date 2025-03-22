package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/db"
	"github.com/oexlkinq/fairdrop/internal/storage"
)

type App struct {
	db       *db.DB
	storage  *storage.Storage
	basePath string
}

func Create(db *db.DB, s *storage.Storage, base string) *App {
	return &App{db, s, base}
}

func (app *App) Run() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	folders := r.Group("/folders")
	app.addFolderRoutes(folders)

	r.Run()
}
