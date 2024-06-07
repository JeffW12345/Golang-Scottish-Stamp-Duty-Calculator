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
        fmt.Printf("TEST PRINT - err value: %s\n", err)
        if err == nil{
            fmt.Println("TEST PRINT - Server GET operation a success")
            break
        }
        if time.Since(startTime) > time.Second {
            fmt.Println("TEST PRINT - panicked")
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