package db

import (
	entity "Concurrencia/internal/app/user/entities"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type UserRepository struct {
	client *dynamodb.DynamoDB
}

func NewUserRepository(client *dynamodb.DynamoDB) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (d *UserRepository) FindUserById(id string) (*entity.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("users_dynamo_db_test"),
		Key: map[string]*dynamodb.AttributeValue{
			"user-id": {
				S: aws.String(id),
			},
		},
	}

	result, err := d.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result != nil {
		nameAttr, found := result.Item["Name"]
		if found {
			name := nameAttr.String()
			return &entity.User{
				Id:   id,
				Name: name,
			}, nil
		}
	}

	return nil, fmt.Errorf("No se encontraron datos en DynamoDB")
}

func (d *UserRepository) CreateUser(user *entity.User) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("users_dynamo_db_test"),
		Item: map[string]*dynamodb.AttributeValue{
			"user-id": {
				S: aws.String(user.Id),
			},
			"Name": {
				S: aws.String(user.Name),
			},
		},
	}
	_, err := d.client.PutItem(input)
	if err != nil {
		fmt.Printf("Error insert user in dynamo %s", err)
		return err
	}
	return nil
}
