package main

import (
	"fmt"
	"strings"
	"time"

	// "example.com/practice/cmdmabager"
	"example.com/practice/filemanager"
	"example.com/practice/prices"
)

const storageDir = "./calculation_results"

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
	
		outputFileName := strings.ToLower(fmt.Sprintf(
			"%s/%s.json",
			storageDir,
			time.Now().Format("20060102_150405"),
		))
		fm := filemanager.New(
			"prices.txt",
			outputFileName,
		)
		// cmdm := cmdmabager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i], errChans[i])

		// if err != nil {
		// 	fmt.Println("Can't process. ", err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <- errChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChans[i]:
			fmt.Println("Done.")
		}
	}
}
