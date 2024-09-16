package prices

import (
	"fmt"

	"example.com/practice/conversion"
	"example.com/practice/filemanager"
)

type TaxIncludedPrices map[string]string

type TaxIncludedPriceJob struct {
	IOManager *filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices TaxIncludedPrices `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

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

	job.IOManager.WriteJSON(job)
}

func NewTaxIncludedPriceJob(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		TaxIncludedPrices: map[string]string{},
		IOManager: fm,
	}
}
