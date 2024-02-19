package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type task struct {
    ID     string  `json:"id"`
    Description  string  `json:"title"`
    Done  bool  `json:"artist"`
}

// albums slice to seed record album data.
var tasks  = []task{
    {ID: "1", Description: "Do laundry", Done: true},
    {ID: "2", Description: "Buy groceries", Done: false},
    {ID: "3", Description: "Maditate by 5 min", Done: true},
}

func main() {
    router := gin.Default()
    router.GET("/tasks", getTasks)
    router.POST("/tasks", addTask)
    router.GET("/tasks/:id", getTask)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getTasks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, tasks)
}


func addTask(c *gin.Context) {
    var newTask task

    if err := c.BindJSON(&newTask); err != nil {
       return
    }
    
    tasks = append(tasks, newTask)

    c.IndentedJSON(http.StatusOK, newTask)
}

func getTask(c *gin.Context) {
    var task task

    for _, t := range tasks {
        if t.ID == c.Param("id") {
            c.IndentedJSON(http.StatusOK, task)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"}) 
}

