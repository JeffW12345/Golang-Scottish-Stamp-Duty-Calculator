package server

import (
	"fmt"
	"os"
)

func CreateValidSortedMockJsonBandConfigFile(path string) {
	validTestData := []byte(`
	[
		{"start": 0, "end": 145000, "percentageTax": 0},
		{"start": 145000.01, "end": 250000, "percentageTax": 0.02},
		{"start": 250000.01, "end": 325000, "percentageTax": 0.05},
		{"start": 325000.01, "end": 750000, "percentageTax": 0.10},
		{"start": 750000.01, "end": 2147483647, "percentageTax": 0.12}
	]`)

	if err := os.WriteFile(path, validTestData, 0644); err != nil {
		panic(fmt.Errorf("failed to create mock json band config file: %v", err))
	}
}

func CreateInvalidMockJsonBandConfigFile(path string) {
	missingCloseBraceData := []byte(`
	[
		{"start": 0, "end": 145000, "percentageTax": 0
	]`)

	if err := os.WriteFile(path, missingCloseBraceData, 0644); err != nil {
		panic(fmt.Errorf("failed to create mock json band config file: %v", err))
	}
}

func CreateEmptyMockJsonBandConfigFile(path string) {
	emptyContent := []byte(``)

	if err := os.WriteFile(path, emptyContent, 0644); err != nil {
		panic(fmt.Errorf("failed to create mock json band config file: %v", err))
	}
}