package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// task represents data about a record task.
// Task represents data that will have title, description and status
type task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Few default Tasks
var tasks = []task{
	{ID: "1", Title: "Workout", Description: "Working out at gym", Status: "pending"},
	{ID: "2", Title: "Study", Description: "Study materials from college", Status: "completed"},
	{ID: "3", Title: "Work", Description: "Part-Time work", Status: "pending"},
}

func main() {
	router := gin.Default()
	// Endpoints start
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTasksByID)
	router.POST("/tasks", postTasks)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	// Endpoints end

	router.Run("localhost:8080")
}

// getTasks responds with the list of all tasks as JSON.
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// postTasks adds an task from JSON received in the request body.
func postTasks(c *gin.Context) {
	var newTask task

	// Call BindJSON to bind the received JSON to
	// newtask.
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return the error message
		return
	}

	// Add the new task to the slice.
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// getTasksByID locates the tasks whose ID value matches the id
// parameter sent by the client, then returns that tasks json as a response.
func getTasksByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of tasks, looking for
	// a task whose ID value matches the parameter.
	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// updateTask updates a task by ID.
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			// Update the task
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status
			c.IndentedJSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// deleteTask deletes a task by ID.
func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task
			c.IndentedJSON(http.StatusNoContent, nil) // Return no content status
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
