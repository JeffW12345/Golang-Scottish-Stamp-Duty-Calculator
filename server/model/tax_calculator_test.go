package model

import (
	"testing"
)

func TestGetTaxDueForPropertyOfValue(t *testing.T) {
	tc := TaxCalculator{}
	tc.addTaxBands(GetValidSortedTestTaxBands())

	t.Run("calculateTaxDue should return error if property value negative number", func(t *testing.T) {
		_, err := tc.calculateTaxDue(-1)
		if err == nil {
			t.Error("calculateTaxDue should return error if property value negative number")
		}
	})
	t.Run("calculateTaxDue should return errors correct error message if property value negative number", func(t *testing.T) {
		_, err := tc.calculateTaxDue(-1)

		want := "value of property cannot be negative"
		if err.Error() != want {
			t.Error("calculateTaxDue returning wrong error message for negative number")
		}
	})

	t.Run("calculateTaxDue should return zero if value of property < highest point on lowest band", func(t *testing.T) {
		val, _ := tc.calculateTaxDue(0)

		var want float64 = 0
		if val != want {
			t.Error("Should have returned zero as property value within first band")
		}
	})

	t.Run("calculateTaxDue should correct amount if price within second band", func(t *testing.T) {
		val, _ := tc.calculateTaxDue(200_000)
		var want float64 = 1100.00
		if val != want {
			t.Error("Should have returned zero as property value within second band")
		}
	})
	t.Run("calculateTaxDue should correct amount if price top of second band", func(t *testing.T) {
		val, _ := tc.calculateTaxDue(250_000)
		var want float64 = 2100.00
		if val != want {
			t.Error("Should have returned zero as property value top of second band")
		}
	})

	t.Run("calculateTaxDue should correct amount if price in top band", func(t *testing.T) {
		val, _ := tc.calculateTaxDue(800_000)
		var want float64 = 54_350.00
		if val != want {
			t.Error("Should have returned zero as property value in top band")
		}
	})
}

func TestAddTaxBands(t *testing.T) {
	tc := TaxCalculator{}
	expectedBands := GetValidSortedTestTaxBands()

	tc.addTaxBands(expectedBands)
	if len(tc.bands) != len(expectedBands) {
		t.Error("Number of tax bands does not match")
		return
	}

	for i, band := range tc.bands {
		if band != expectedBands[i] {
			t.Errorf("Tax band at index %d does not match, expected: %v, got: %v", i, expectedBands[i], band)
		}
	}
}
