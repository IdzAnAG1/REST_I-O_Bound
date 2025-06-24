package server

import (
	"github.com/gin-gonic/gin"
	"rest_io_bound/internal/handlers"
)

type WebServer struct {
	port string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{port: port}
}

func (ws *WebServer) Launch() {
	router := gin.Default()

	router.Handle("GET", "/", handlers.GetTasks)
	router.Handle("POST", "/", handlers.CreateTask)
	router.Handle("GET", "/:id", handlers.GetTaskById)
	router.Handle("DELETE", "/:id", handlers.DeleteTaskByID)

	err := router.Run(":" + ws.port)
	if err != nil {
		return
	}
}
