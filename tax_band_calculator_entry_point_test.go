package main

import (
	"fmt"
	"testing"
)

type MockTaxCalculatorNoErrorReturned struct {
	propertyValue float32
}

func (o *MockTaxCalculatorNoErrorReturned) getTaxDueForPropertyOfValue() (float32, error) {
	return 0, nil
}

type MockTaxCalculatorErrorReturned struct {
	propertyValue float32
}

func (o *MockTaxCalculatorErrorReturned) getTaxDueForPropertyOfValue() (float32, error) {
	return 0, fmt.Errorf("mock error")
}

func TestIsServerReadyYet(t *testing.T) {
	t.Run("isServerReadyYet should return true if no error returned from API", func(t *testing.T) {
		tc := &MockTaxCalculatorNoErrorReturned{propertyValue: 200_000}
		got := isServerReadyYet(tc)
		want := true
		if got != want {
			t.Error("isServerReadyYet should return true if no error returned from API but did not")
		}
	})

	t.Run("isServerReadyYet should panic if server not ready after 2 seconds", func(t *testing.T) {
		done := make(chan bool, 1)
		message := make(chan string, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					message <- "OK"
					done <- true
				} else {
					message <- "error"
					done <- false
				}
			}()

			fmt.Println("Starting test for panic")
			tc := &MockTaxCalculatorErrorReturned{propertyValue: 200_000}
			isServerReadyYet(tc)
		}()

		result := <-done
		panicMsg := <-message

		fmt.Println("Outcome of panic test: ", panicMsg)

		if !result {
			t.Error("isServerReadyYet should panic if server not ready after 2 seconds")
		}
	})
}
