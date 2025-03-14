package routes

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/utils"
)

// TODO: #2 это значение надо бы передавать ещё и в форму пользователю, чтобы браузер знал что сервер не примет данных больше этого значения
// 32MiB
const maxBunchFilesUploadSize = 32 << 20

func (app *App) addFolderRoutes(folders *gin.RouterGroup) {
	folders.POST("/create", app.createFolder)
	folders.GET("/:password", app.listFolder)
	folders.POST("/:password", app.uploadFile)
	folders.DELETE("/:password", app.deleteFolder)
	folders.GET("/:password/:filename", app.serveFile)
}

func (app *App) createFolder(c *gin.Context) {
	var pw string
	for {
		pw = utils.GeneratePassword()

		if !app.db.TestFolder(c, pw) {
			break
		}
	}

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
	for i, filename := range *files {
		(*files)[i] = path.Join(app.basePath, filename)
	}

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

	// #2 TODO: чекнуть, реально ли это нужно
	if strings.TrimPrefix(params.Filename, "/") == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// satinized filename
	filename := path.Base(params.Filename)
	fullpath := filepath.Join(app.storage.Path, params.Password, filename)

	c.File(fullpath)
}

func (app *App) uploadFile(c *gin.Context) {
	var params FolderParams

	err := c.BindUri(&params)
	if err != nil {
		return
	}

	if !app.db.TestFolder(c, params.Password) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = c.Request.ParseMultipartForm(maxBunchFilesUploadSize)
	if err != nil {
		log.Println(fmt.Errorf("upload file: parse multipart form: %w", err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for filename, files := range c.Request.MultipartForm.File {
		// #3 TODO: в files могут лежать несколько файлов с одинаковым name из формы. мб здесь стоит делать чтото большее чем просто отбрасывать их
		file := files[0]
		fullpath := filepath.Join(app.storage.Path, params.Password, path.Base(filename))

		err := c.SaveUploadedFile(file, fullpath)
		if err != nil {
			// #3 TODO: мб стоит удалять уже сохранёные файлы при ошибке
			panic(fmt.Errorf("save file from form: %w", err))
		}
	}
}

func (app *App) deleteFolder(c *gin.Context) {
	var params FolderParams

	err := c.BindUri(&params)
	if err != nil {
		return
	}

	if !app.db.DeleteFolder(c, params.Password) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	app.storage.DeleteFolder(params.Password)
}
