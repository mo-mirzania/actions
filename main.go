package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type req struct {
	Name string `json:"name"`
}

type res struct {
	Greet string `json:"greet"`
}

func main() {
	lambda.Start(handler)
}

func handler(req req) (res, error) {
	return res{
		Greet: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}
