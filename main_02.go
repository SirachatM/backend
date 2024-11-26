package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  //Employee API Method
  router.GET("/employee", getEmployee)
  router.POST("/employee", postEmployee)
  router.PUT("/employee", putEmployee)
  router.DELETE("/employee", deleteEmployee)
    
  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee GET Method!",
  })
}

func postEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee POST Method!",
  })
}

func putEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee PUT Method!",
  })
}

func deleteEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee DELETE Method!",
  })
}