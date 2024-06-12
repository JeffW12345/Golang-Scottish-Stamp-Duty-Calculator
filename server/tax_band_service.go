package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type TaxBand struct {
	Start         float32 `json:"start"`
	End           float32 `json:"end"`
	PercentageTax float32 `json:"percentageTax"`
}

type TaxBands struct {
	Bands []TaxBand
	JsonConfigFilePath string
}

func (tbs *TaxBands) ImportAndProcessTaxBands() {
	tbs.importTaxBands()
	tbs.sortByStartingValue()
	tbs.performDataIntegrityCheckOnBands()
}

func (tbs *TaxBands) getTaxBands() []TaxBand {
	return tbs.Bands
}

func (tbs *TaxBands) importTaxBands() {
	file := openFile(tbs)
	defer file.Close()

	fileContents := jsonBandsConfigFileAsString(file)

	if err := json.Unmarshal(fileContents, &tbs.Bands); err != nil {
			panic(fmt.Errorf("json conversion error: %v", err))
		}
	
	//TODO - Test this
	if len(tbs.Bands) == 0 {
		panic("no bands imported from config file")
	}
}

func openFile(tbs *TaxBands) *os.File {
	file, err := os.Open(tbs.JsonConfigFilePath)
	if err != nil {
		panic(fmt.Errorf("file open error: %v", err))
	}
	return file
}

func jsonBandsConfigFileAsString(file *os.File) []byte {
	fileContents, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("file read error: %v", err))
	}
	return fileContents
}

func (tbs *TaxBands) sortByStartingValue() {
	sort.Slice(tbs.Bands, func(i, j int) bool {
		return tbs.Bands[i].Start < tbs.Bands[j].Start
	  })
}

func (tbs *TaxBands) performDataIntegrityCheckOnBands() {
	// Check if has at least one band
	//TODO - Check if sorted by start value
	// Check if each start value is less than every end value
	//TODO - Check if start, end or percentageTax negative
}