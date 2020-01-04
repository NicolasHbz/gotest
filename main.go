package main

import (
	usersController "gotest/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1RouterGroup := router.Group("/api/v1")

	usersController.InitController(v1RouterGroup)

	router.Run(":3002")
}
