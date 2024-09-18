package handler

import (
	"context"
	"go-lambda-starter/bootstrap"
	"go-lambda-starter/src/types"
	"log"
)

// StartEventProcess will execute the business logic
func StartEventProcess(ctx context.Context, request types.EventPayload) error {
	log.Printf("RECEIVED request : %+v", request)

	// TODO: your business logic starts here
	app := bootstrap.NewApplication()
	// ...
	// Dummy logging, you can remove this later
	for key := range app.Config.Data {
		log.Printf("env key %v = %v", key, app.Config.GetString(key))
	}
	return nil
}
