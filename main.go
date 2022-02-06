package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/danesparza/tztest/data"
)

// Message is a custom struct event type to handle the Lambda input
type Message struct {
	Timezone string `json:"timezone"`
}

// HandleRequest handles the AWS lambda request
func HandleRequest(ctx context.Context, msg Message) (string, error) {

	if err := data.SetTimezone(msg.Timezone); err != nil {
		log.Fatal(err) // most likely timezone not loaded in Docker OS
	}
	t := data.GetTime(time.Now())
	response := t.String()

	//	Return our response
	return response, nil
}

func main() {
	//	Immediately forward to Lambda
	lambda.Start(HandleRequest)
}
