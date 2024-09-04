package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := getBlance()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		// panic("Can't proceed.")
	}

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
			writeBalance(accountBalance)
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
			writeBalance(accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 4:
			fmt.Println("Bye bye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func  writeBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBlance() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("failed to read balance")
	}

	balanceText, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("failed to parse stored balance")
	}

	return balanceText, nil
}
