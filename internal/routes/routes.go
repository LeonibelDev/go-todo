package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leonibeldev/go-todo/internal/database"
	model "github.com/leonibeldev/go-todo/internal/models"
)

var DbExist = database.VerifyDBexist()
var db = database.DbConnection()

// GetAllTasks function
func GetAllTasks(c *gin.Context) {

	query := `
		SELECT * FROM Tasks;
	`

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Datecreated)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to scan task"})
			fmt.Println(err.Error())
			return
		}

		tasks = append(tasks, task)
	}

	// check if tasks is nil
	if tasks == nil {
		tasks = []model.Task{}

	}

	// response tasks
	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

// GetTask (only one task)
func GetTask(c *gin.Context) {
	id := c.Param("id")

	query := `
		SELECT * FROM Tasks WHERE ID = ?;
	`

	rows, err := db.Query(query, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to query database",
		})

		return
	}

	defer rows.Close()

	if !rows.Next() {
		c.JSON(500, gin.H{
			"message": "tasks not found",
		})
		return
	}

	var task model.Task
	if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Datecreated); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}

// CreateTask function
func CreateTask(c *gin.Context) {
	var t model.Task
	c.BindJSON(&t)

	query := `
		INSERT INTO Tasks(title, description) VALUES(?, ?);
	`

	result, err := db.Exec(query, t.Title, t.Description)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// check if task was created
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(500, gin.H{
			"message": "failed to create task",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "task created",
	})
}

// UpdateTask function
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var t model.Task
	c.BindJSON(&t)

	query := `
		UPDATE Tasks SET title = ?, description = ? WHERE ID = ?;
	`

	result, err := db.Exec(query, t.Title, t.Description, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// check if task was updated
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(500, gin.H{
			"message": "task not found",
		})
		return
	}

	// if task was updated
	c.JSON(200, gin.H{
		"message": "task updated",
	})
}

// DeleteTask function
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	query := `
		DELETE FROM Tasks WHERE ID = ?;
	`

	result, err := db.Exec(query, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// check if task was deleted
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(500, gin.H{
			"message": "task not found",
		})
		return
	}

	// if task was deleted
	c.JSON(200, gin.H{
		"message": "task deleted",
	})
}
