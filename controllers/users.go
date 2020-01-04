package userscontroller

import (
	"gotest/middlewares"
	"gotest/models"
	usersvalidator "gotest/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitController(router *gin.RouterGroup) {
	usersRouter := router.Group("/users")
	{
		usersRouter.GET("/", getUsers)
		usersRouter.POST("/", middlewares.Validate(middlewares.ValidationObject{
			Body: &usersvalidator.PostUserBody{},
		}), postUser)
	}
}

func getUsers(c *gin.Context) {
	users := []models.User{
		models.User{Firstname: "nico"},
		models.User{Firstname: "nico2", Lastname: "H"},
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func postUser(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}
