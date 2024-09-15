package prices

import (
	"fmt"
	"strings"
	"time"

	"example.com/practice/conversion"
	"example.com/practice/filemanager"
)

const storageDir = "./calculation_results"

type TaxIncludedPrices map[string]string

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices TaxIncludedPrices
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println("File read error: ", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Converting string to float failed. ", err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)

		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fileName := strings.ToLower(fmt.Sprintf(
		"%s/%s.json",
		storageDir,
		time.Now().Format("20060102_150405"),
	))

	filemanager.WriteJSON(fileName, job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		TaxIncludedPrices: map[string]string{},
	}
}
