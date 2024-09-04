package main

import (
	"fmt"
)

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Print(`What do you want to do?
1. Check balance
2. Deposit money
3. Withdraw money
4. Exit
`)

		var choice int
		fmt.Print("Please select prefered option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance: ", accountBalance)	
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount should be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("Not enough funds!")
				continue
			} 

			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount should be greater than 0.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 4:
			fmt.Println("Bye bye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}