package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGet displays application index page
func IndexGet(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
