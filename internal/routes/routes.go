package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/db"
	"github.com/oexlkinq/fairdrop/internal/storage"
)

type App struct {
	db      *db.DB
	storage *storage.Storage
}

func Create(db *db.DB, s *storage.Storage) *App {
	return &App{db, s}
}

func (app *App) Run() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(func(ctx *gin.Context) {
		// #2 TODO: сменить * на домен или пару доменов
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	serveFrontend := makeServeFrontend()
	r.Use(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/folders/") {
			return
		}

		serveFrontend(ctx)
	})

	folders := r.Group("/folders")
	app.addFolderRoutes(folders)

	r.Run()
}
