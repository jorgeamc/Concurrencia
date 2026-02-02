package main

import (
	"Concurrencia/internal/app/user/repository/db"
	"Concurrencia/internal/app/user/service"
	"Concurrencia/web"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

type dependencies struct {
	service web.UserService
	queue   web.MsgService
}

func main() {

	dep := dependencies{}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Cambia esto a tu región preferida
	}))
	r := gin.Default()
	dynamoClient := dynamodb.New(sess)
	userRepository := db.NewUserRepository(dynamoClient)
	dep.service = service.NewService(userRepository)
	r.POST("/create-user", web.CreateUser(dep.service))
	r.GET("/get-user", web.GetUser(dep.service))
	r.POST("/queue/publish", web.PublishTopic(dep.queue))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
