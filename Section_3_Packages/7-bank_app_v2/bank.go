package main

import (
	"fmt"

	"example.com/bank_app_v2/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetValueFromFile(accountBalanceFile, 10000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		// return or panic("Can't continue")
	}

	fmt.Println("------ Welcome To The Bank ------")
	for {
		fmt.Println("What do you wish to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is:", accountBalance)
		
		case 2:
			var depositAmount float64
			fmt.Print("Amount to be deposited: ")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("Updated balance:", accountBalance)
			fileops.WriteValueToFile(accountBalance)

		case 3:
			var withdrawAmount float64
			fmt.Print("Amount to be deducted: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Incorrect value! Value should be more than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Incorrect value! Insufficient Balance.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Updated balance:", accountBalance)
			fileops.WriteValueToFile(accountBalance)

		default:
			fmt.Println("Exiting...")
			fmt.Print("Thank you for using our app!")
			return
		}
	}
}