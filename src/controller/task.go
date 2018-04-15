package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/andoshin11/go-web-app/src/model"
	"github.com/gin-gonic/gin"
)

// TasksGET returns list of task
func TasksGET(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tasks := []model.Task{}

	// iterate result
	for result.Next() {
		task := model.Task{}
		var id uint
		var createdAt, updatedAt time.Time
		var title string

		err = result.Scan(&id, &createdAt, &updatedAt, &title)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		tasks = append(tasks, task)
	}
	fmt.Println(tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// TaskPOST creates a new task
func TaskPOST(c *gin.Context) {
	db := model.DBConnect()

	title := c.PostForm("title")
	now := time.Now()

	task := &model.Task{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := task.Save(db)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. title: %s", title)
}

// TaskPATCH updates a task
func TaskPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	task, err := model.TaskByID(db, uint(id))
	if err != nil {
		panic(err.Error())
	}

	title := c.PostForm("title")
	now := time.Now()

	task.Title = title
	task.UpdatedAt = now

	err = task.Update(db)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(task)
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// TaskDELETE deletes a task
func TaskDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	// Check if record exists
	task, err := model.TaskByID(db, uint(id))
	if err != nil {
		panic(err.Error())
	}

	err = task.Delete(db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")
}
