package main

import (
	"github.com/andoshin11/go-web-app/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGet)
	router.Run(":8080")
}
