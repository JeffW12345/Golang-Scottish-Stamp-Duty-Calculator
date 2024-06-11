package server

type TaxBand struct {
	start         float32
	end           float32
	percentageTax float32
}

type TaxBands struct {
	bands []TaxBand
}

func importAndProcessTaxBands() {
	importTaxBands()
	sortBandsByStartingValue()
	performDataIntegrityCheckOnBands()
}

func getTaxBands() ([] TaxBand) {
	//TODO 
	return []TaxBand{}
}

func importTaxBands() {
	//TODO
}

func sortBandsByStartingValue() {
	//TODO 
}

func performDataIntegrityCheckOnBands() {
	//TODO
}

func (tb *TaxBands) addTaxBand(start, end, percentageTax float32) {
	band := TaxBand{
		start:         start,
		end:           end,
		percentageTax: percentageTax,
	}
	tb.bands = append(tb.bands, band)
}