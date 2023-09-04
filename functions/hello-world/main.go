package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	type APIResponse struct {
		Message string `json:"message"`
	}

	response := APIResponse{
		Message: "Hello World!",
	}

	responseInJSON, err := json.Marshal(response)
	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error",
		}, err
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       string(responseInJSON),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
