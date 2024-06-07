package main

import (
	"fmt"
	"testing"

	"bou.ke/monkey"
)

func didPanicHappen(f func()) (didPanic bool) {
    defer func() {
        if r := recover(); r != nil {
            didPanic = true
        }
    }()
    f()
    return
}

func TestUpdateChannelWhenServerReady(t *testing.T) {
    defer monkey.UnpatchAll()

    t.Run("Should update channel if no error returned from API", func(t *testing.T) {
        mockIsServerReady := make(chan bool)
        monkey.Patch(getTaxDueForPropertyOfValue, func(valueOfProperty float32) (float32, error) {
            return 0, nil
        })
        go updateChannelWhenServerReady(mockIsServerReady)
        value := <-mockIsServerReady
        if !value {
            t.Error("channel not updated when getTaxDueForPropertyOfValue returns no error")
        }
    })

    t.Run("Should panic if server not ready after 1 second", func(t *testing.T) {
        mockIsServerReady := make(chan bool)
		
		monkey.UnpatchAll()
		monkey.Patch(getTaxDueForPropertyOfValue, func(valueOfProperty float32) (float32, error) {
			return 0, fmt.Errorf("mock error")
		})

        didPanic := didPanicHappen(func() {
            updateChannelWhenServerReady(mockIsServerReady)
        })

        if !didPanic {
            t.Error("expected panic but did not happen")
        }
    })
}
