package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	prompt("Revenue: ", &revenue)
	prompt("Expenses: ", &expenses)
	prompt("Tax Rate: ", &taxRate)

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f \n", ebt)
	fmt.Printf("Profit: %.2f \n", profit)
	fmt.Printf("Ratio: %.2f \n", ratio)
}

func prompt(requestText string, val *float64) {
	fmt.Print(requestText)
	fmt.Scan(val)
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}
