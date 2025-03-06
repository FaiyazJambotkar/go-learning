package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profitAfterTax := earningsBeforeTax * (1 - taxRate / 100)
	ratio := earningsBeforeTax / profitAfterTax

	formattedEBT := fmt.Sprintf("Earnings Before Tax: %v\n", earningsBeforeTax)
	formattedProfit := fmt.Sprintf("Profit After Tax: %v\n", profitAfterTax)

	fmt.Print(formattedEBT, formattedProfit)
	fmt.Printf("Ratio: %.2f", ratio)
}