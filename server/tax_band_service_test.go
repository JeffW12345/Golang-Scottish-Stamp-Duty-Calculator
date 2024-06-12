package server

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
