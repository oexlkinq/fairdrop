package routes

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/db"
	"github.com/oexlkinq/fairdrop/internal/storage"
)

type App struct {
	db       *db.DB
	storage  *storage.Storage
	basePath string
}

func Create(db *db.DB, s *storage.Storage, basePath string) *App {
	return &App{db, s, basePath}
}

func (app *App) Run() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(func(ctx *gin.Context) {
		// #2 TODO: сменить * на домен или пару доменов
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	rootGroup := r.Group(app.basePath)

	serveFrontend := makeServeFrontend(app.basePath)
	apiPrefix := path.Join(app.basePath, "/folders/")
	r.Use(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, apiPrefix) {
			return
		}

		serveFrontend(ctx)
	})

	folders := rootGroup.Group("/folders")
	app.addFolderRoutes(folders)

	r.Run()
}
