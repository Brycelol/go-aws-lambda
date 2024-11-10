package helloperson

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

// Person Struct encapsulating first and last name
type Person struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(event json.RawMessage) error {
	var person Person

	err := json.Unmarshal(event, &person)
	if err != nil {
		return err
	}

	log.Printf("Hello %s %s\n", *person.FirstName, *person.LastName)
	return nil
}
