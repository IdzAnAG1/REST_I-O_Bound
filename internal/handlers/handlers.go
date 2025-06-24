package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rest_io_bound/internal/models"
	"rest_io_bound/internal/variables"
	"time"
)

type AutoProcess interface {
	LaunchAutoProcess()
}

func GetTasks(c *gin.Context) {
	c.JSON(200, variables.Storage)
}

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task.Created = time.Now()
	task.Status = "Created"
	key, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task.LaunchAutoProcess()

	variables.Storage[key.String()] = &task
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task := variables.Storage[id]
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
	c.JSON(200, task)
}

func DeleteTaskByID(c *gin.Context) {
	ID := c.Param("id")
	delete(variables.Storage, ID)
	c.JSON(200, gin.H{"message": "task deleted"})
}
