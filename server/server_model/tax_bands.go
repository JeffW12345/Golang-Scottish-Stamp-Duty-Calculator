package server_model

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type TaxBands struct {
	Bands              []TaxBand
	JsonConfigFilePath string
}

func (tbs *TaxBands) ImportAndProcessTaxBands() {
	tbs.importTaxBands()
	tbs.sortByStartingValue()
	tbs.performDataIntegrityCheckOnBands()
}

func (tbs *TaxBands) importTaxBands() {
	file := openFile(tbs)
	defer file.Close()

	fileContents := jsonBandsConfigFileAsString(file)

	if err := json.Unmarshal(fileContents, &tbs.Bands); err != nil {
		panic(fmt.Sprintf("json conversion error: %v", err))
	}

	if len(tbs.Bands) == 0 {
		panic("no bands imported from config file")
	}
}

func openFile(tbs *TaxBands) *os.File {
	file, err := os.Open(tbs.JsonConfigFilePath)
	if err != nil {
		panic(fmt.Sprintf("file open error: %v", err))
	}
	return file
}

func jsonBandsConfigFileAsString(file *os.File) []byte {
	fileContents, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("file read error: %v", err))
	}
	return fileContents
}

func (tbs *TaxBands) sortByStartingValue() {
	sort.Slice(tbs.Bands, func(i, j int) bool {
		return tbs.Bands[i].Start < tbs.Bands[j].Start
	})
}

func (tbs *TaxBands) performDataIntegrityCheckOnBands() {
	var previousEnd float64 = -1

	for i, v := range tbs.Bands {
		// Check if end value is greater than start value
		if v.Start >= v.End {
			panic(fmt.Sprintf("Band start values must be less than end values: index %d, start %f, end %f", i, v.Start, v.End))
		}

		// Check if current band start value is greater than previous end value
		if i > 0 && v.Start <= previousEnd {
			panic(fmt.Sprintf("Band start value must be greater than the previous band's end value: index %d, start %f, previous end %f", i, v.Start, previousEnd))
		}

		// Check that the band variables are all positive numbers
		if v.Start < 0 || v.End < 0 || v.PercentageTax < 0 {
			panic(fmt.Sprintf("Band variables must be positive numbers: index %d, start %f, end %f, percentageTax %f", i, v.Start, v.End, v.PercentageTax))
		}

		// Update previous end value for the next iteration
		previousEnd = v.End
	}
}
