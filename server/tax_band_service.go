package server

import (
	"encoding/json"
	"io"
	"os"
	"fmt"
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
	tbs.sortBandsByStartingValue()
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

func (tbs *TaxBands) sortBandsByStartingValue() {
	//TODO
}

func (tbs *TaxBands) performDataIntegrityCheckOnBands() {
	//TODO
}

func (tb *TaxBands) addTaxBand(taxBand TaxBand) {
	tb.Bands = append(tb.Bands, taxBand)
}