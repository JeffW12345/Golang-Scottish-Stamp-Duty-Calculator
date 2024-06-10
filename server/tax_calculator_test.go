package server

import (
	"testing"
)


func TestGetTaxDueForPropertyOfValue(t *testing.T) {
	t.Run("getTaxDueForPropertyOfValue should return error if property value negative number", func(t *testing.T) {
		_, err := calculateTaxDue(-1)
		if err == nil {
			t.Error("getTaxDueForPropertyOfValue should return error if property value negative number")
		}
	})
	t.Run("getTaxDueForPropertyOfValue should return errors correct error message if property value negative number", func(t *testing.T) {
		_, err := calculateTaxDue(-1)

		want := "value of property cannot be negative"
		if err.Error() != want {
			t.Error("getTaxDueForPropertyOfValue returning wrong error message for negative number")
		}
	})
}