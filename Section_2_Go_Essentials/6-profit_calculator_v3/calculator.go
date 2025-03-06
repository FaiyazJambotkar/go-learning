package main

import (
	"errors"
	"fmt"
	"os"
)

func getUserInput(text string) (float64, error){
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("invalid value \nenter value above 0")
	}

	return input, nil
}

func calculateValues(revenue, expenses, taxRate float64) (earningsBeforeTax, profitAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profitAfterTax = earningsBeforeTax * (1 - taxRate / 100)
	ratio = earningsBeforeTax / profitAfterTax
	
	return earningsBeforeTax, profitAfterTax, ratio
}

func storeValuesInFile(earningsBeforeTax, profitAfterTax, ratio float64) {
	writeToFile := fmt.Sprintf("Earnings Before Tax: %.1f\nProfit After Tax: %.1f\nRatio: %.3f\n", earningsBeforeTax, profitAfterTax, ratio)
	os.WriteFile("Calculated_values.txt", []byte(writeToFile), 0644)
}

func printOutput(text string, value float64) {
	fmt.Printf(text, value)
}

func main() {
	revenue, err := getUserInput("Enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	taxRate, err := getUserInput("Enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profitAfterTax, ratio := calculateValues(revenue, expenses, taxRate)

	storeValuesInFile(earningsBeforeTax, profitAfterTax, ratio)

	printOutput("Earnings Before Tax: %.1f\n", earningsBeforeTax)
	printOutput("Profit After Tax: %.1f\n", profitAfterTax)
	printOutput("Ratio: %.3f\n", ratio)
}