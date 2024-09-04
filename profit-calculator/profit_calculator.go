package main

import (
	"fmt"
	"errors"
	"os"
	"time"
)

func main() {
	var revenue, expenses, taxRate float64
	var err error

	prompt("Revenue: ", &revenue)
	err = validateInput(revenue)
	if err != nil {
		renderError(err)
		return
	}
	prompt("Expenses: ", &expenses)
	err = validateInput(expenses)
	if err != nil {
		renderError(err)
		return
	}
	prompt("Tax Rate: ", &taxRate)
	err = validateInput((taxRate))
	if err != nil {
		renderError(err)
		return
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	result := fmt.Sprintf(`Earnings Before Tax: %.2f
Profit: %.2f
Ratio: %.2f
`, ebt, profit, ratio)

	fmt.Print(result)
	err = storeCalculationResult(revenue, expenses, taxRate, result)
	if err != nil {
		fmt.Println("_____ERROR_____")
		panic(err)
	}
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

func validateInput(input float64) error {
	if input <= 0 {
		return errors.New("provided value should b greater than zero")
	}

	return nil
}

func renderError(errorText error) {
	fmt.Println("-----ERROR-----")
	fmt.Println(errorText)
	fmt.Println("-----ERROR-----")
}

func storeCalculationResult(revenue float64, expenses float64, taxRate float64, result string) error {
	dataToStore := fmt.Sprintf(`Revenue: %.2f Expenses: %.2f Tax Rate: %.2f
%s`, revenue, expenses, taxRate, result)
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("calculation_result_%s.txt", timestamp)

	err := os.WriteFile(filename, []byte(dataToStore), 0644)

	return err
}