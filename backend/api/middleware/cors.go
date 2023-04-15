package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware provide cors middleware
var CorsMiddleware gin.HandlerFunc

func init() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE", "PATCH"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Accept"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowAllOrigins = true
	CorsMiddleware = cors.New(corsConfig)
}