package main

import (
	"github.com/andoshin11/go-web-app/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", controller.TasksGET)
		v1.POST("/tasks", controller.TaskPOST)
	}

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
