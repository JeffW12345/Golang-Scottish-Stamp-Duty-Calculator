package main

import (
	"time"

	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/server"
)

type TaxRetrievalInterface interface {
	getTaxDue() (float32, error)
}

type TaxRetrieval struct {
	propertyValue float32
}

func (o *TaxRetrieval) getTaxDue() (float32, error) {
	// TODO: Function body
	return 0, nil
}

func displayTaxDueForProperty(valueOfProperty float32) {
	// TODO: Function body
}

func isServerReadyYet(taxRetriever TaxRetrievalInterface) bool {
	startTime := time.Now()
	for {
		if time.Since(startTime) > (time.Second * 2) {
			panic("Server timed out")
		}
		_, err := taxRetriever.getTaxDue()
		if err == nil {
			break
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
	return true
}

func waitTillServerReady() {
	tc := &TaxRetrieval{propertyValue: 0}
	for !isServerReadyYet(tc) {
		continue
	}
}

func main() {
	tbs := server.TaxBands{}
	tbs.JsonConfigFilePath = "resources/tax_band_configuration.json"
	tbs.ImportAndProcessTaxBands()
	go server.ServerSetup()

	waitTillServerReady()

	displayTaxDueForProperty(200_000)
}