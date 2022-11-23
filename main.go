package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type RequestEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, req RequestEvent) (string, error) {
	return fmt.Sprintf("Hello %s", req.Name), nil

}

func main() {
	lambda.Start(HandleRequest)
}
