package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) AddFolderRoutes(rg *gin.RouterGroup) {
	folders := rg.Group("/folders")

	folders.POST("/", app.createFolder)
}
func (app *App) createFolder(c *gin.Context) {
	var params struct {
		PwComplexity PasswordComplexity `form:"password-complexity"`
	}
	err := c.ShouldBind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "payload": params})
		return
	}

	folder, err := app.db.CreateFolder(c, params.PwComplexity, c.ClientIP())
	if err != nil {
		log.Println(fmt.Errorf("create folder handler: %w", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, folder)
}
