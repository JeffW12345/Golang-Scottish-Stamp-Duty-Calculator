package server

import (
	"os"
	"testing"
)

func TestImportTaxBands(t *testing.T) {
	path := "mock_band_json_file.json"
	defer os.Remove(path)

	tbs := &TaxBands{
		JsonConfigFilePath: path,
	}
	CreateMockJsonBandConfigFile(path)
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