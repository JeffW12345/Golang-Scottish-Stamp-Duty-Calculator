package server

import (
	"fmt"
	"os"
)

func GetValidSortedTestTaxBands() []TaxBand {
	return []TaxBand{
		{Start: 0, End: 145000, PercentageTax: 0},
		{Start: 145000.01, End: 250000, PercentageTax: 0.02},
		{Start: 250000.01, End: 325000, PercentageTax: 0.05},
		{Start: 325000.01, End: 750000, PercentageTax: 0.10},
		{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
	}
}

func GetValidUnsortedTestTaxBands() []TaxBand {
	return []TaxBand{
		{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
		{Start: 0, End: 145000, PercentageTax: 0},
		{Start: 250000.01, End: 325000, PercentageTax: 0.05},
		{Start: 325000.01, End: 750000, PercentageTax: 0.10},
		{Start: 145000.01, End: 250000, PercentageTax: 0.02},
	}
}

func CreateValidSortedMockJsonBandConfigFile(path string) {
	testData := []byte(`
	[
		{"start": 0, "end": 145000, "percentageTax": 0},
		{"start": 145000.01, "end": 250000, "percentageTax": 0.02},
		{"start": 250000.01, "end": 325000, "percentageTax": 0.05},
		{"start": 325000.01, "end": 750000, "percentageTax": 0.10},
		{"start": 750000.01, "end": 2147483647, "percentageTax": 0.12}
	]`)

	if err := os.WriteFile(path, testData, 0644); err != nil {
		panic(fmt.Errorf("failed to create mock json band config file: %v", err))
	}
}

func CreateEmptyMockJsonBandConfigFile(path string) {
	testData := []byte(``)

	if err := os.WriteFile(path, testData, 0644); err != nil {
		panic(fmt.Errorf("failed to create mock json band config file: %v", err))
	}
}

func CreateUnsortedMockJsonBandConfigFile(path string) {
	testData := []byte(`
	[
		{"start": 0, "end": 145000, "percentageTax": 0},
		{"start": 325000.01, "end": 750000, "percentageTax": 0.10},
		{"start": 750000.01, "end": 2147483647, "percentageTax": 0.12},
		{"start": 145000.01, "end": 250000, "percentageTax": 0.02},
		{"start": 250000.01, "end": 325000, "percentageTax": 0.05}
	]`)

	if err := os.WriteFile(path, testData, 0644); err != nil {
		panic(fmt.Errorf("failed to create mock json band config file: %v", err))
	}
}
