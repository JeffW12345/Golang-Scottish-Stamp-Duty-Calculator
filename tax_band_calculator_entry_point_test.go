package main

import (
	"testing"

	"bou.ke/monkey"
)

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

	t.Run("Should update channel if no error returned from API", func(t *testing.T) {
		mockIsServerReady := make(chan bool)
		monkey.Patch(getTaxDueForPropertyOfValue, func() (float32, error) {return 0, nil})
		updateChannelWhenServerReady(mockIsServerReady)
		channelContent, channelUpdated := <-mockIsServerReady

		if !channelUpdated || !channelContent {
			t.Error("channel not updated when getTaxDueForPropertyOfValue returns an error")
		}
	})
}