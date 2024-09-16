package main

import (
	"fmt"
	"strings"
	"time"

	"example.com/practice/filemanager"
	"example.com/practice/prices"
)

const storageDir = "./calculation_results"

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		outputFileName := strings.ToLower(fmt.Sprintf(
			"%s/%s.json",
			storageDir,
			time.Now().Format("20060102_150405"),
		))
		fm := filemanager.New(
			"prices.txt",
			outputFileName,
		)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
