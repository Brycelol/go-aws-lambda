package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	lambda.Start(handleMessage)
}

func handleMessage(event events.SQSEvent) error {
	log.Printf("Received %d SQS events!\n", len(event.Records))

	for _, record := range event.Records {
		log.Printf("Received SNS message [%s]\n", record.MessageId)
	}

	return nil
}
