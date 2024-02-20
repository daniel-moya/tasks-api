package main

import (
  "database/sql"
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  _ "github.com/lib/pq"
)

type task struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Done  bool  `json:"artist"`
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
    router.PUT("/tasks/:id", func (c *gin.Context) { updateTask(c, db) })
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

    rows, err := db.Query("SELECT * FROM tasks;")
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var cTask task
        rows.Scan(&cTask.ID, &cTask.Title, &cTask.Done)
        tasks = append(tasks, cTask)
    }

    c.IndentedJSON(http.StatusOK, tasks)
}


func createTask(c *gin.Context, db *sql.DB) {
    var newTask task

    if err := c.BindJSON(&newTask); err != nil {
       return
    }
    err := db.QueryRow("INSERT INTO tasks (title, done) VALUES($1, $2) RETURNING id", newTask.Title, newTask.Done).Scan(&newTask.ID)
	if err != nil {
		log.Fatal(err)
	}
    
    c.IndentedJSON(http.StatusOK, newTask)
}

func readTask(c *gin.Context, db *sql.DB) {
    rows, err := db.Query("SELECT * FROM tasks WHERE id = ($1)", c.Param("id"))

    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var cTask task
         rows.Scan(&cTask.ID, &cTask.Title, &cTask.Done)
        c.IndentedJSON(http.StatusOK, cTask)
        return
    }

    c.IndentedJSON(http.StatusNotFound, "Task not found")
}

func updateTask(c *gin.Context, db *sql.DB) {

    c.IndentedJSON(http.StatusOK, "TO BE DONE")
}

func removeTask(c *gin.Context, db *sql.DB) {
    c.IndentedJSON(http.StatusOK, "TO BE DONE")
}


func generateTasksReport(c *gin.Context, db *sql.DB) {

    c.IndentedJSON(http.StatusOK, "TO BE DONE")
}

