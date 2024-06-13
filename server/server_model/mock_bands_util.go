package server_model

func GetValidSortedTestTaxBands() []TaxBand {
	return []TaxBand{
		{Start: 0, End: 145000, PercentageTax: 0},
		{Start: 145000.01, End: 250000, PercentageTax: 0.02},
		{Start: 250000.01, End: 325000, PercentageTax: 0.05},
		{Start: 325000.01, End: 750000, PercentageTax: 0.10},
		{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
	}
}

func GetValidButUnsortedTestTaxBands() []TaxBand {
	return []TaxBand{
		{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
		{Start: 0, End: 145000, PercentageTax: 0},
		{Start: 250000.01, End: 325000, PercentageTax: 0.05},
		{Start: 325000.01, End: 750000, PercentageTax: 0.10},
		{Start: 145000.01, End: 250000, PercentageTax: 0.02},
	}
}
