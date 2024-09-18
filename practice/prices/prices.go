package prices

import (
	"fmt"

	"example.com/practice/conversion"
	"example.com/practice/iomanager"
)

type TaxIncludedPrices map[string]string

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices TaxIncludedPrices `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadData()

	if err != nil {
		errChan <- err
		doneChan <- false
		return
	}

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)

		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.IOManager.WriteResult(job)

	doneChan <- true
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		TaxIncludedPrices: map[string]string{},
		IOManager: fm,
	}
}
