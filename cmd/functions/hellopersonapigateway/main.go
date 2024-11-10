package hellopersonapigateway

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Go Lambda for integration with API Gateway

// Person Struct encapsulating first and last name
type Person struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type ResponseBody struct {
	Message *string `json:"message"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var person Person

	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	msg := fmt.Sprintf("Hello there %s %s", *person.FirstName, *person.LastName)
	responseBody := ResponseBody{
		Message: &msg,
	}

	jsonBytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonBytes),
	}

	return response, nil
}
