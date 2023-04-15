package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOk response ok
func GetOk(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// TBD remove it
func TBD(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Not Implemented Yet")
}
