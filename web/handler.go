package web

import (
	entity "Concurrencia/internal/app/user/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Id   string `json:"id" binding:"required"`
			Name string `json:"name" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user := entity.User{
			Id:   input.Id,
			Name: input.Name,
		}

		err := service.CreateUser(&user)

		if err != nil {
			fmt.Printf("Error creating user: %s", err)
			c.JSON(http.StatusFailedDependency, err)
		}

		c.JSON(http.StatusCreated, "User Created")
	}

}

func GetUser(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.DefaultQuery("user", "")
		getUser, err := service.GetUser(user)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}
		c.JSON(http.StatusOK, getUser)
	}
}

func PublishTopic(service MsgService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Msg string `json:"message" binding:"required"`
		}
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		msg := entity.Msg{
			Value: input.Msg,
		}

		fmt.Println("El mensaje es: ", msg.Value)

		//service.PublishTopic(&msg)
		c.JSON(http.StatusOK, "Publish result OK")
	}
}
