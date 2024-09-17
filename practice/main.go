package main

import (
	"fmt"

	"example.com/practice/cmdmabager"
	// "example.com/practice/filemanager"
	"example.com/practice/prices"
)

const storageDir = "./calculation_results"

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// outputFileName := strings.ToLower(fmt.Sprintf(
		// 	"%s/%s.json",
		// 	storageDir,
		// 	time.Now().Format("20060102_150405"),
		// ))
		// fm := filemanager.New(
		// 	"prices.txt",
		// 	outputFileName,
		// )
		cmdm := cmdmabager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Can't process. ", err)
		}
	}
}
