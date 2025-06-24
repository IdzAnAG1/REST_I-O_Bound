package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rest_io_bound/internal/models"
	"time"
)

func GetTasks(c *gin.Context) {
	c.JSON(200, models.Storage)
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
	models.Storage[key.String()] = &task
}
