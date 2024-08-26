package main

import (
	"context"
	"os"

	"football_tracker/src/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) error {
	footballService := services.NewFootballService()

	footballService.BrocastRecentCompletedMatches()

	return nil
}

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "test_function" {
		lambda.Start(handler)
	} else {
		handler(context.Background())
	}
}
