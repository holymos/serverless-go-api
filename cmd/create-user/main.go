package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/holymos/serverless-go-api/internal/user"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var usr user.User

	err := json.Unmarshal([]byte(request.Body), &usr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	usr.Id = uuid.New().String()
	usr.CreatedAt = time.Now().String()
	// usr.UpdatedAt = time.Now().String()

	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv(user.UsersTable)),
		Item:      user.FormatUserToDynamo(usr),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadGateway,
		}, nil
	}

	rsp, err := json.Marshal(usr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadGateway,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body: string(rsp),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
