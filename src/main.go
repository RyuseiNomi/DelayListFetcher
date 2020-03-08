package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {
	log.Print("Hoge")
}

func main() {
	lambda.Start(handler)
}
