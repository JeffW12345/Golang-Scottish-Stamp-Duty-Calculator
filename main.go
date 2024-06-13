package main

import (
	"log"

	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/client/client_model"
	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/server"
)

func main() {
	go server.ServerSetup()

	client := &client_model.TaxRequest{}
	client.WaitTillServerReady()
	log.Println("Server ready")

	// PASS PROPERTY VALUE INTO THIS METHOD: 
	client.DisplayTaxDueForProperty(200_000)
}
