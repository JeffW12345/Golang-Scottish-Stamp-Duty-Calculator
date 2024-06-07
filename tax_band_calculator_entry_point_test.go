package main

import (
	"testing"
	"time"

	"bou.ke/monkey"
)

func didPanicHappen() bool {	
	return recover() != nil
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
			t.Error("channel not updated when getTaxDueForPropertyOfValue returns an error")
		}
	})

	t.Run("Should panic if server not ready after 1 second", func(t *testing.T) {
		mockIsServerReady := make(chan bool)
		monkey.Patch(getTaxDueForPropertyOfValue, func(valueOfProperty float32) (float32, error) {
			return 0, nil
		})
		go func() {
			startTime := time.Now()
			for time.Since(startTime) < 2000{
			}
			if !didPanicHappen() {
				t.Error("channel not updated when getTaxDueForPropertyOfValue returns an error")
			}
		} ()
		go updateChannelWhenServerReady(mockIsServerReady)
	})
}