package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/utils"
)

func (app *App) addFolderRoutes(folders *gin.RouterGroup) {
	folders.POST("/create", app.createFolder)
	folders.GET("/:password", app.listFolder)
	folders.GET("/:password/:filename", app.serveFile)
}

func (app *App) createFolder(c *gin.Context) {
	pw := utils.GeneratePassword()
	// TODO: проверять уникальность пароля по базе

	app.storage.CreateFolder(pw)

	folder := app.db.CreateFolder(c, pw, c.ClientIP())

	c.JSON(http.StatusOK, folder)
}

type FolderParams struct {
	Password string `uri:"password" binding:"required"`
}

func (app *App) listFolder(c *gin.Context) {
	var params FolderParams

	err := c.BindUri(&params)
	if err != nil {
		return
	}
	pw := params.Password

	if !app.db.TestFolder(c, pw) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	files := app.storage.ListFolder(pw)

	c.JSON(http.StatusOK, files)
}

type FileParams struct {
	FolderParams
	Filename string `uri:"filename" binding:"required"`
}

func (app *App) serveFile(c *gin.Context) {
	var params FileParams

	err := c.BindUri(&params)
	if err != nil {
		return
	}

	if strings.TrimPrefix(params.Filename, "/") == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.File(app.storage.GetFilepath(params.Password, params.Filename))
}
