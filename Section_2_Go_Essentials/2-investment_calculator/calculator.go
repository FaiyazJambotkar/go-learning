package main

import (
	"fmt"
	"math"
)

func main() {
	const infaltionRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("For Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+infaltionRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}