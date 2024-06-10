package main

import (
	"time"

	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/server"
)

type TaxCalculationInterface interface {
	getTaxDueForPropertyOfValue() (float32, error)
}

type TaxCalculator struct {
	propertyValue float32
}

func (o *TaxCalculator) getTaxDueForPropertyOfValue() (float32, error) {
	// Simulating a real calculation or request
	return 0, nil
}

func displayTaxDueForProperty(valueOfProperty float32) {
	// TODO: Function body
}

func isServerReadyYet(taxCalculator TaxCalculationInterface) bool {
	startTime := time.Now()
	for {
		if time.Since(startTime) > (time.Second * 2) {
			panic("Server timed out")
		}
		_, err := taxCalculator.getTaxDueForPropertyOfValue()
		if err == nil {
			break
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
	return true
}

func getTaxDueForPropertyOfValue(valueOfProperty float32) (float32, error) {
	// Simulating a real calculation or request
	return 0, nil
}

func main() {
	go server.ServerSetup()

	waitTillServerReady()
	
	displayTaxDueForProperty(200_000)
}

func waitTillServerReady() {
	tc := &TaxCalculator{propertyValue: 200_000}
	for !isServerReadyYet(tc) {
		continue
	}
}
