package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.GET("/get", get)
	r.Run()
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"maker": "ok", "machine-number": "test", "high-w": 10, "low-w": 1})
}
