package main

import (
	"StoriTxChallenge/internal/application/service"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading receiver.env file: %v", err)
	}
	service.BalanceGenerator("./resources/txns.csv")
}
