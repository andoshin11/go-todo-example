package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGET displays application index page
func IndexGET(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
