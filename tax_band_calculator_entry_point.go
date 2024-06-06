package main

import (
	"tax-calculator/server"
	"time"
)

func updateChannelWhenServerReady(isServerReady chan bool) {
    var hasPingedSuccessfully bool
    var hasTimedOut bool
    startTime := time.Now()
    for (!hasPingedSuccessfully && !hasTimedOut) {
        _ , err := getTaxDueForPropertyOfValue(200_000)
        if err == nil{
            hasPingedSuccessfully = true
        }
        if time.Since(startTime) > 1000 {
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