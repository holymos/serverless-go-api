package user

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const UsersTable = "USERS_TABLE"

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func FormatUserToDynamo(usr User) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(usr.Id),
		},
		"firstName": {
			S: aws.String(usr.FirstName),
		},
		"lastName": {
			S: aws.String(usr.LastName),
		},
		"email": {
			S: aws.String(usr.Email),
		},
		"username": {
			S: aws.String(usr.Username),
		},
	}
}
