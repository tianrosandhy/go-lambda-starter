package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go-lambda-starter/src/handler"
	"go-lambda-starter/src/requestparser"
	"go-lambda-starter/src/types"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tianrosandhy/goconfigloader"
)

func main() {
	cfg := goconfigloader.NewConfigLoader()

	asLocal := cfg.GetBool("RUN_AS_LOCAL")
	if asLocal {
		LocalHandler()
	} else {
		lambda.Start(EventLambdaHandler)
	}
}

func LocalHandler() {
	log.Printf("RUNNING in LocalHandler")
	ctx := context.Background()
	request := types.EventPayload{}
	mockEvent := "./mock-event/request.json"
	mockData, err := os.ReadFile(mockEvent)
	if err != nil {
		log.Printf("FAIL to Read Mock Event Data : %+v", err)
	}
	json.Unmarshal(mockData, &request)
	err = handler.StartEventProcess(ctx, request)

	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("DONE")
	}
}

func EventLambdaHandler(ctx context.Context, request map[string]interface{}) error {
	log.Printf("RUNNING in EventLambdaHandler")
	payload := types.EventPayload{}
	// will try to convert either RabbitMQ, HTTP Gateway, or Raw Event to correct payload
	requestparser.ParseEventRequest(request, &payload)

	err := handler.StartEventProcess(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
