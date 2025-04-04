package routes

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oexlkinq/fairdrop/frontend"
)

func makeServeFrontend() gin.HandlerFunc {
	devFrURLString := os.Getenv("DEV_FRONTEND_URL")

	if devFrURLString == "" {
		return static.Serve("/", static.EmbedFolder(frontend.FrontendContentFS, "dist"))
	} else {
		log.Println("using proxy for frontend")

		devFrURL, err := url.Parse(devFrURLString)
		if err != nil {
			panic(fmt.Errorf("parse frontend url: %w", err))
		}

		rp := httputil.NewSingleHostReverseProxy(devFrURL)

		return func(ctx *gin.Context) {
			rp.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}
