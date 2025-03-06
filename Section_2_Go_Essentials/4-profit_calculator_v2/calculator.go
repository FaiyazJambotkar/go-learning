package main

import (
	"fmt"
)

func main() {
	revenue := setValue("Enter Revenue: ")
	expenses := setValue("Enter Expenses: ")
	taxRate := setValue("Enter Tax Rate: ")

	earningsBeforeTax, profitAfterTax, ratio := calculateValues(revenue, expenses, taxRate)

	printOutput("Earnings Before Tax: %.1f\n", earningsBeforeTax)
	printOutput("Profit After Tax: %.1f\n", profitAfterTax)
	printOutput("Ratio: %.3f\n", ratio)
}

func setValue(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculateValues(revenue, expenses, taxRate float64) (earningsBeforeTax, profitAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profitAfterTax = earningsBeforeTax * (1 - taxRate / 100)
	ratio = earningsBeforeTax / profitAfterTax
	
	return earningsBeforeTax, profitAfterTax, ratio
}

func printOutput(text string, value float64) {
	fmt.Printf(text, value)
}
