package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leonibeldev/go-todo/internal/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r := gin.Default()

	// get all tasks
	r.GET("/api/tasks", routes.GetAllTasks)

	// create task
	r.POST("/api/tasks", routes.CreateTask)

	// get one task
	r.GET("/api/tasks/:id", routes.GetTask)

	// update task
	r.PUT("/api/tasks/:id", routes.UpdateTask)

	// delete task
	r.DELETE("/api/tasks/:id", routes.DeleteTask)

	// run server
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	r.Run(":" + PORT)
}
