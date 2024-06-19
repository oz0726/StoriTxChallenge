package main

import (
	fileReader "StoriTxChallenge/internal/infrastructure/ports/input"
	"fmt"
)

func main() {
	fmt.Println(fileReader.ReadFile("./resources/txns.csv"))
}
