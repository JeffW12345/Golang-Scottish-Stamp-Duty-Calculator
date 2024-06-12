package model

import (
	"os"
	"testing"
	"log"
)

func TestImportTaxBands(t *testing.T) {
	t.Run("importTaxBands works as expected when importing valid tax bands", func(t *testing.T) {
		// Create mock json config file containing valid sorted tax bands
		path := "mock_band_valid_json_file.json"
		CreateValidSortedMockJsonBandConfigFile(path)

		// Delete mock config file when test complete
		defer os.Remove(path)

		// Import bands from dummy config file and convert to TaxBands object.
		taxBands := &TaxBands{
			JsonConfigFilePath: path,
		}
		taxBands.importTaxBands()
		got := taxBands.Bands

		want := GetValidSortedTestTaxBands()
		// Does the TaxBands object contain the correct number of tax bands?
		if len(got) != len(want) {
			t.Errorf("Number of imported tax bands not as expected. Got %d, expected %d", len(got), len(want))
			return
		}
		// Is each tax band as expected?
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Tax band at index %d does not match, expected: %v, got: %v", i, got[i], want[i])
			}
		}
	})

	t.Run("importTaxBands panics if no tax bands created", func(t *testing.T) {
		// Create mock empty band config file
		path := "mock_band_empty_json_file.json"
		CreateEmptyMockJsonBandConfigFile(path)

		// Delete dummy config file when test complete
		defer os.Remove(path)

		panicHappened := make(chan bool, 1)
		message := make(chan string, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					message <- "test passes as panic happened"
					panicHappened <- true
				} else {
					message <- "test fails as panic did not happen"
					panicHappened <- false
				}
			}()
			// Import empty config file and attempt toconvert to TaxBands object.
			taxBands := &TaxBands{
				JsonConfigFilePath: path,
			}
			taxBands.importTaxBands()
		}()

		log.Println("Outcome of panic test: ", <-message)

		if !<-panicHappened {
			t.Error("importTaxBands should panic as config file empty")
		}
	})

	t.Run("importTaxBands panics if object creation error", func(t *testing.T) {
		// Create mock invalid band config file
		path := "mock_band_invalid_json_file.json"
		CreateInvalidMockJsonBandConfigFile(path)

		// Delete dummy config file when test complete
		defer os.Remove(path)

		panicHappened := make(chan bool, 1)
		message := make(chan string, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					message <- "test passes as panic happened"
					panicHappened <- true
				} else {
					message <- "test fails as panic did not happen"
					panicHappened <- false
				}
			}()
			// Import invalid config file and attempt toconvert to TaxBands object.
			taxBands := &TaxBands{
				JsonConfigFilePath: path,
			}
			taxBands.importTaxBands()
		}()

		log.Println("Outcome of panic test: ", <-message)

		if !<-panicHappened {
			t.Error("importTaxBands should panic as config file not valid")
		}
	})
}

func TestSortByStartingValue(t *testing.T) {
	tbs := TaxBands{}
	tbs.Bands = GetValidButUnsortedTestTaxBands()
	tbs.sortByStartingValue()
	got := tbs.Bands

	want := GetValidSortedTestTaxBands()

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Tax band at index %d does not match", i)
		}
	}
}

func TestPerformDataIntegrityCheckOnBands(t *testing.T) {
	tests := []struct {
		name     string
		taxBands TaxBands
		wantPanic bool
	}{
		{
			name: "Valid Bands",
			taxBands: TaxBands{
				Bands: []TaxBand{
					{Start: 0, End: 145000, PercentageTax: 0},
					{Start: 145000.01, End: 250000, PercentageTax: 0.02},
					{Start: 250000.01, End: 325000, PercentageTax: 0.05},
					{Start: 325000.01, End: 750000, PercentageTax: 0.10},
					{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
				},
			},
			wantPanic: false,
		},
		{
			name: "Invalid Start and End Values",
			taxBands: TaxBands{
				Bands: []TaxBand{
					{Start: 0, End: 145000, PercentageTax: 0},
					{Start: 145000.01, End: 250000, PercentageTax: 0.02},
					{Start: 250000.01, End: 250000.01, PercentageTax: 0.05},
					{Start: 325000.01, End: 750000, PercentageTax: 0.10},
					{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
				},
			},
			wantPanic: true,
		},
		{
			name: "Non-increasing Start Values",
			taxBands: TaxBands{
				Bands: []TaxBand{
					{Start: 0, End: 145000, PercentageTax: 0},
					{Start: 145000.01, End: 250000, PercentageTax: 0.02},
					{Start: 200000.01, End: 325000, PercentageTax: 0.05},
					{Start: 325000.01, End: 750000, PercentageTax: 0.10},
					{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
				},
			},
			wantPanic: true,
		},
		{
			name: "Negative Values",
			taxBands: TaxBands{
				Bands: []TaxBand{
					{Start: -1, End: 145000, PercentageTax: 0},
					{Start: 145000.01, End: 250000, PercentageTax: 0.02},
					{Start: 250000.01, End: 325000, PercentageTax: 0.05},
					{Start: 325000.01, End: 750000, PercentageTax: 0.10},
					{Start: 750000.01, End: 2147483647, PercentageTax: 0.12},
				},
			},
			wantPanic: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !test.wantPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				} else {
					if test.wantPanic {
						t.Errorf("expected panic but did not happen")
					}
				}
			}()
			test.taxBands.PerformDataIntegrityCheckOnBands()
		})
	}
}