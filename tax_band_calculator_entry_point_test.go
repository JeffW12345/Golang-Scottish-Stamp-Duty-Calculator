package main

import (
	"testing"

	"bou.ke/monkey"
)

func didPanicHappen() bool {	
	return recover() != nil
}

func mockUpdateChannelWhenServerReady(isServerReady chan bool) {
	
	go updateChannelWhenServerReady(isServerReady)

}

func UpdateChannelWhenServerReadyTest(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("Should update channel if no error returned from API", func(t *testing.T) {
		mockIsServerReady := make(chan bool)
		monkey.Patch(getTaxDueForPropertyOfValue, func() (float32, error) {return 0, nil})
		updateChannelWhenServerReady(mockIsServerReady)
		channelContent, channelUpdated := <-mockIsServerReady

		if !channelUpdated || !channelContent {
			t.Error("channel not updated when getTaxDueForPropertyOfValue returns an error")
		}
	})

	t.Run("Should panic if server not ready after 1 second", func(t *testing.T) {
		mockIsServerReady := make(chan bool)
		monkey.Patch(getTaxDueForPropertyOfValue, func() (float32, error) {return 0, nil})
		updateChannelWhenServerReady(mockIsServerReady)
		channelContent, channelUpdated := <-mockIsServerReady

		if !channelUpdated || !channelContent {
			t.Error("channel not updated when getTaxDueForPropertyOfValue returns an error")
		}
	})
}