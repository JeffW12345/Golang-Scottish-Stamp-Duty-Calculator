package model

import (
	"fmt"
	"math"
)

type TaxCalculator struct {
	bands []TaxBand
}

func (tc *TaxCalculator) CalculateTaxDue(valueOfProperty float64) (float64, error) {
	if valueOfProperty < 0 {
		return 0, fmt.Errorf("value of property cannot be negative")
	}

	var taxDue float64
	taxDue = 0.0
	for _, band := range tc.bands {
		if valueOfProperty <= band.Start {
			break
		}
		if valueOfProperty >= band.End {
			taxDue += ((band.End - band.Start) * band.PercentageTax)
		} else {
			taxDue += ((valueOfProperty - band.Start) * band.PercentageTax)
			break
		}
	}

	taxDueRounded := float64(math.Round(float64(taxDue)*100) / 100)
	return taxDueRounded, nil
}

func (tc *TaxCalculator) AddTaxBands(taxBands []TaxBand) {
	tc.bands = append(tc.bands, taxBands...)
}
