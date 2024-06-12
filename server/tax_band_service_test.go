package server

import (
	"os"
	"testing"
)

func TestImportTaxBands(t *testing.T) {

	t.Run("importTaxBands works as expected when importing valid tax bands", func(t *testing.T) {
		// Create mock json config file containing valid sorted tax bands
		path := "mock_band_json_file.json"
		CreateValidSortedMockJsonBandConfigFile(path)

		// Delete mock config file when test complete
		defer os.Remove(path)
	
		// Import bands from dummy config file and convert to TaxBands object.
		taxBands := &TaxBands{
			JsonConfigFilePath: path,
		}
		taxBands.importTaxBands()
		got := taxBands.Bands

		// Does the TaxBands object contain the correct number of tax bands? 
		if len(got) != 5 {
			t.Error("Number of tax bands does not match")
			return
		}

		// Is each tax band as expected?
		want := GetValidSortedTestTaxBands()
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Tax band at index %d does not match, expected: %v, got: %v", i, got[i], want[i])
			}
		}
	})
}

func TestSortByStartingValue(t *testing.T) {
	tbs := TaxBands{}
	tbs.Bands = GetUnsortedTestTaxBands()
	tbs.sortByStartingValue()

	want := GetValidSortedTestTaxBands()

	// Check if the lengths of both slices are the same
	if len(tbs.Bands) != len(want) {
		t.Error("Number of tax bands does not match")
		return
	}

	// Compare each tax band to see if they are equal
	for i := range want {
		if tbs.Bands[i] != want[i] {
			t.Errorf("Tax band at index %d does not match", i)
		}
	}
}

func TestImportTaxBandsEmptyConfigFile(t *testing.T) {
	//TODO - Complete
	path := "mock_band_json_file.json"
	defer os.Remove(path)

	tbs := &TaxBands{
		JsonConfigFilePath: path,
	}
	CreateValidSortedMockJsonBandConfigFile(path)
	tbs.importTaxBands()

	expectedBands := tbs.Bands

	if len(expectedBands) != 5 {
		t.Error("Number of tax bands does not match")
		return
	}

	for i, band := range expectedBands {
		if band != expectedBands[i] {
			t.Errorf("Tax band at index %d does not match, expected: %v, got: %v", i, expectedBands[i], band)
		}
	}
}
