package main

import (
	"github.com/RyuseiNomi/delay_reporter_lm/src/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
