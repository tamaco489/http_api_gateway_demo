package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	log.Println("[INFO] 1. Request", request)

	log.Println("[INFO] 2. Request Path:", request.Path)
	log.Println("[INFO] 2. Request HTTPMethod:", request.HTTPMethod)
	log.Println("[INFO] 2. Request Headers:", request.Headers)
	log.Println("[INFO] 2. Request PathParameters:", request.PathParameters)
	log.Println("[INFO] 2. Request QueryStringParameters:", request.QueryStringParameters)
	log.Println("[INFO] 2. Request Body:", request.Body)

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
