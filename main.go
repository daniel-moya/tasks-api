package main

import (
  "database/sql"
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  _ "github.com/lib/pq"
)

type task struct {
    ID     int32  `json:"id"`
    Title  string  `json:"title"`
    Done  bool  `json:"done"`
}

var connectionStr string = "postgresql://postgres:dev@127.0.0.1/postgres?sslmode=disable"

func main() {
    db := initDB()
    router := gin.Default()
    
    // Collection routes
    router.GET("/tasks", func (c *gin.Context) { getTasks(c, db) })
    
    // CRUD for tasks
    router.POST("/tasks", func (c *gin.Context) { createTask(c, db) })
    router.GET("/tasks/:id", func (c *gin.Context) { readTask(c, db) })
    router.PUT("/tasks", func (c *gin.Context) { updateTask(c, db) })
    router.DELETE("/tasks/:id", func (c *gin.Context) { removeTask(c, db) })

    // Report routes
    router.GET("/tasks/report", func (c *gin.Context) { generateTasksReport(c, db) })

    router.Run("localhost:8080")
}

func initDB() *sql.DB {
  db, err := sql.Open("postgres", connectionStr)
  if err != nil {
    log.Fatal(err)
  }
  return db
}

// getAlbums responds with the list of all albums as JSON.
func getTasks(c *gin.Context, db *sql.DB) {
    tasks := []task{}

    rows, err := db.Query("SELECT id, title, done FROM tasks;")

    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tasks"})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var cTask task
        rows.Scan(&cTask.ID, &cTask.Title, &cTask.Done)
        tasks = append(tasks, cTask)
    }

    // Check for errors during row iteration
    if rowsErr := rows.Err(); rowsErr != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while iterating over tasks"})
        return
    }

    c.IndentedJSON(http.StatusOK, tasks)
}


func createTask(c *gin.Context, db *sql.DB) {
    var newTask task

    if err := c.BindJSON(&newTask); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON payload"})
        return
    }
    _, err := db.Exec("INSERT INTO tasks (title, done) VALUES($1, $2)", newTask.Title, newTask.Done)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new Task"})
        return
	}
    
    c.IndentedJSON(http.StatusOK, newTask)
}

func readTask(c *gin.Context, db *sql.DB) {
    var cTask task
    err := db.QueryRow("SELECT id, title, done FROM tasks WHERE id = ($1)", c.Param("id")).Scan(&cTask.ID, &cTask.Title, &cTask.Done)

    if err != nil {
        if err == sql.ErrNoRows {
            c.IndentedJSON(http.StatusNotFound, "Task not found")
            return
        }

        c.IndentedJSON(http.StatusInternalServerError, "Error fetching task")
        return
    }

    c.IndentedJSON(http.StatusOK, cTask)
}

func updateTask(c *gin.Context, db *sql.DB) {
    var uTask task

    if err := c.BindJSON(&uTask); err != nil {
        log.Fatal(err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON payload"})
        return
    }
    _, err := db.Exec("UPDATE tasks SET title = ($1), done = ($2) WHERE id = ($3);",
        &uTask.Title,
        &uTask.Done,
        &uTask.ID)

    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed updating task"})
		return
	}

     // Return success response
    c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func removeTask(c *gin.Context, db *sql.DB) {
    _, err := db.Exec("DELETE FROM tasks WHERE id = ($1)", c.Param("id"))

    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed deleting task"})
		return
	}

     // Return success response
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func generateTasksReport(c *gin.Context, db *sql.DB) {

    c.IndentedJSON(http.StatusOK, "TO BE DONE")
}

