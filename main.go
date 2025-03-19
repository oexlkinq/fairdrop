package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	db *DB
}

func main() {
	config := Populate()

	app := App{
		db: Connect(config.DBFile),
	}

	r := gin.Default()
	app.AddFolderRoutes(&r.RouterGroup)
	r.Run()

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
