package main

import (
	"math"
	"fmt"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount,expectedReturnRate, years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future value(including inflation): %.2f", futureRealValue)

	// fmt.Println("Future value:", futureValue)
	// fmt.Println(formattedFV, formattedRFV)
	fmt.Printf(`Future value: %.2f
Future value(including inflation): %.2f
`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return
}