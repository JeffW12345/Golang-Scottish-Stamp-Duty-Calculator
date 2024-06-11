package server

import (
	"os"
	"testing"
)

func TestImportTaxBandsHappyPath(t *testing.T) {
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

func TestSortByStartingValue(t *testing.T) {
    tbs := TaxBands{}
    tbs.Bands = GetValidUnsortedTestTaxBands()
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
