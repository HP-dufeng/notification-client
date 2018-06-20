package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "./wwwroot")

	router.Run("localhost:8080")
}
