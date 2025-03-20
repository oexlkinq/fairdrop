package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/utils"
)

func (app *App) addFolderRoutes(folders *gin.RouterGroup) {
	folders.POST("/create", app.createFolder)
	folders.GET("/:password", app.listFolder)
}

func (app *App) createFolder(c *gin.Context) {
	pw := utils.GeneratePassword()
	// TODO: проверять уникальность пароля по базе

	app.storage.CreateFolder(pw)

	folder := app.db.CreateFolder(c, pw, c.ClientIP())

	c.JSON(http.StatusOK, folder)
}

func (app *App) listFolder(c *gin.Context) {
	var params struct {
		Password string `uri:"password" binding:"required"`
	}

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
