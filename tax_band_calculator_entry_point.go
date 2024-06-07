package main

import (
	"tax-calculator/server"
	"time"
    "fmt"
)

func updateChannelWhenServerReady(isServerReady chan bool) {
    startTime := time.Now()
    for (true) {
        _ , err := getTaxDueForPropertyOfValue(200_000)
        if err == nil{
            break
        }
        if time.Since(startTime) > time.Second {
            panic("Server timed out")
        }
    }
    isServerReady <- true
}

func displayTaxDueForProperty(valueOfProperty float32) {
    // TODO: Function body
}

func getTaxDueForPropertyOfValue(valueOfProperty float32) (float32, error) {
    /*
    TODO - Write code to return the following: 

    - The value of the tax (or zero if the server was unable to process the request) 

    - Nil (or Error object if the server was unable to process the request)
    */
    return 0, nil
}

func main() {
    isServerReady := make(chan bool)
    go server.ServerSetup()
    go updateChannelWhenServerReady(isServerReady)
    
    <- isServerReady
    displayTaxDueForProperty(200_000)
}