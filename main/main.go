package main

import (
	"StoriTxChallenge/internal/application/service"
	"github.com/joho/godotenv"
	"log"
)
import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(runLambda)
}

func runLambda() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading receiver.env file: %v", err)
	}
	service.BalanceGenerator("txns.csv")
}
