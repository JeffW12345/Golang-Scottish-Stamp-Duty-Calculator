package client_model

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// MockTaxRetrieverNoErrorReturned is a mock implementation of TaxRetrievalInterface that returns no error
type MockTaxRetrieverNoErrorReturned struct {
	propertyValue float32
}

func (o *MockTaxRetrieverNoErrorReturned) getTaxDue() (float32, error) {
	return 0, nil
}

func (o *MockTaxRetrieverNoErrorReturned) displayTaxDueForProperty(valueOfProperty float32) {
	// Mock implementation
}

func (o *MockTaxRetrieverNoErrorReturned) isServerReadyYet() bool {
	startTime := time.Now()
	for {
		if time.Since(startTime) > (time.Second * 2) {
			panic("Server timed out")
		}
		_, err := o.getTaxDue()
		if err == nil {
			break
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
	return true
}

// MockTaxRetrieverErrorReturned is a mock implementation of TaxRetrievalInterface that returns an error
type MockTaxRetrieverErrorReturned struct {
	propertyValue float32
}

func (o *MockTaxRetrieverErrorReturned) getTaxDue() (float32, error) {
	return 0, fmt.Errorf("mock error")
}

func (o *MockTaxRetrieverErrorReturned) displayTaxDueForProperty(valueOfProperty float32) {
	// Mock implementation
}

func (o *MockTaxRetrieverErrorReturned) isServerReadyYet() bool {
	startTime := time.Now()
	for {
		if time.Since(startTime) > (time.Second * 2) {
			panic("Server timed out")
		}
		_, err := o.getTaxDue()
		if err == nil {
			break
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
	return true
}

// TestIsServerReadyYet tests the isServerReadyYet function
func TestIsServerReadyYet(t *testing.T) {
	t.Run("isServerReadyYet should return true if no error returned from API", func(t *testing.T) {
		tc := &MockTaxRetrieverNoErrorReturned{propertyValue: 0}
		got := tc.isServerReadyYet()
		want := true
		if got != want {
			t.Error("isServerReadyYet should return true if no error returned from API but did not")
		}
	})

	t.Run("isServerReadyYet should panic if server not ready after 2 seconds", func(t *testing.T) {
		panicHappened := make(chan bool, 1)
		message := make(chan string, 1)

		checkIfPanicsIfTimeout(message, panicHappened)

		log.Println("Outcome of panic test: ", <-message)

		if !<-panicHappened {
			t.Error("isServerReadyYet should panic if server not ready after 2 seconds")
		}
	})
}

// checkIfPanicsIfTimeout checks if the function panics on timeout
func checkIfPanicsIfTimeout(message chan string, done chan bool) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				message <- "test passes as panic happened"
				done <- true
			} else {
				message <- "test fails as panic did not happen"
				done <- false
			}
		}()
		tc := &MockTaxRetrieverErrorReturned{propertyValue: 200000}
		tc.isServerReadyYet()
	}()
}
