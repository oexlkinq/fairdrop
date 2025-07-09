package routes

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/internal/db"
	"github.com/oexlkinq/fairdrop/internal/utils"
)

// TODO: #2 это значение надо бы передавать ещё и в форму пользователю, чтобы браузер знал что сервер не примет данных больше этого значения
// 32MiB
const maxBunchFilesUploadSize = 32 << 20

func (app *App) addFolderRoutes(folders *gin.RouterGroup) {
	folders.POST("/create", app.createFolderHandler)
	folders.POST("/push", app.pushFilesHandler)
	folders.GET("/:password", app.listFolderHandler)
	folders.POST("/:password", app.uploadFilesHandler)
	folders.DELETE("/:password", app.deleteFolderHandler)
	folders.GET("/:password/:filename", app.serveFileHandler)
}

func (app *App) createFolderHandler(c *gin.Context) {
	folder := app.createFolder(c)

	c.JSON(http.StatusOK, folder)
}

func (app *App) createFolder(c *gin.Context) *db.Folder {
	var pw string
	for {
		pw = utils.GeneratePassword()

		if !app.db.TestFolder(c, pw) {
			break
		}
	}

	app.storage.CreateFolder(pw)

	return app.db.CreateFolder(c, pw, c.ClientIP())
}

type FolderParams struct {
	Password string `uri:"password" binding:"required"`
}

func (app *App) listFolderHandler(c *gin.Context) {
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

func (app *App) serveFileHandler(c *gin.Context) {
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

// аплоад файлов в существующую папку
func (app *App) uploadFilesHandler(c *gin.Context) {
	var params FolderParams

	err := c.BindUri(&params)
	if err != nil {
		return
	}

	// чекнуть есть ли папка в базе
	if !app.db.TestFolder(c, params.Password) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	app.uploadFiles(c, params.Password)
}

func (app *App) uploadFiles(c *gin.Context, password string) {
	// прочитать форму. может вернуть ошибку превышения размера
	err := c.Request.ParseMultipartForm(maxBunchFilesUploadSize)
	if err != nil {
		log.Println(fmt.Errorf("upload file: parse multipart form: %w", err))
		c.AbortWithStatus(http.StatusRequestEntityTooLarge)
		return
	}

	// перебор файлов формы. здесь может быть несколько ключей потому что в форме ключи для файлов могут повторяться
	files := c.Request.MultipartForm.File["files"]
	for _, file := range files {
		fullpath := filepath.Join(app.storage.Path, password, path.Base(file.Filename))

		err := c.SaveUploadedFile(file, fullpath)
		if err != nil {
			// #3 TODO: мб стоит удалять уже сохранёные файлы при ошибке
			panic(fmt.Errorf("save file from form: %w", err))
		}
	}
}

// аплоад с созданием папки
func (app *App) pushFilesHandler(c *gin.Context) {
	folder := app.createFolder(c)
	app.uploadFiles(c, folder.Password)

	c.JSON(http.StatusOK, folder)
}

func (app *App) deleteFolderHandler(c *gin.Context) {
	var params FolderParams

	err := c.BindUri(&params)
	if err != nil {
		return
	}

	if !app.db.DeleteFolder(c, params.Password) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	app.storage.DeleteFolder(params.Password)
}
