package main

import (
	"log"

	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/client/client_model"
	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/server"
)

func main() {
	go server.ServerSetup()

	client := &client_model.TaxRequest{}
	client.PropertyValue = 100
	client.WaitTillServerReady() // Times out if not ready after 2 seconds.
	log.Println("\n\nCLICK HERE: http://localhost:8080/\n\n")

	for {} // Infinite loop to prevent the program terminating.
}
