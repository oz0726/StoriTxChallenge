package main

import (
	"StoriTxChallenge/internal/application/service"
)

func main() {
	service.BalanceGenerator("./resources/txns.csv")
}
