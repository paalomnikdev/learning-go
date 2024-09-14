package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPrices map[string]float64

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices TaxIncludedPrices
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	
	if err != nil {
		fmt.Println("Can't open file. ", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Scanner error. ", err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for i, v := range lines {
		floatPrice, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Println("Converting string to float failed. ", err)
			file.Close()
			return
		}

		prices[i] = floatPrice
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		price, err := strconv.ParseFloat(
			fmt.Sprintf("%.2f", taxIncludedPrice),
			64,
		)

		if err != nil {
			fmt.Println("Can't load data. ", err)
			return
		}

		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		TaxIncludedPrices: map[string]float64{},
	}
}
