package main

import (
	"context"
	"log"
	"strconv"

	"date-diff-api/internal/datediff"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type QueryParams struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Units string `json:"units"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse query parameters from API Gateway event
	params := QueryParams{
		Start: request.QueryStringParameters["start"],
		End:   request.QueryStringParameters["end"],
		Units: request.QueryStringParameters["units"],
	}

	input, err := datediff.ParseInput(params.Start, params.End, params.Units)
	if err != nil {
		log.Printf("Failed to parse input: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	diff := datediff.CalculateDateDiff(input)

	log.Printf(
		"The difference between %s and %s is %.1d %s.\n",
		input.Start.Format("Jan 2, 2006"),
		input.End.Format("Jan 2, 2006"),
		diff,
		input.Units,
	)

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		Body:            strconv.Itoa(diff),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
