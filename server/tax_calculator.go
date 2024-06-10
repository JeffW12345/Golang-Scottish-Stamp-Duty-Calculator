package server

import (
	"fmt"
)

func calculateTaxDue(valueOfProperty float32) (float32, error) {
	if valueOfProperty < 0 {
		return 0, fmt.Errorf("value of property cannot be negative")
	}

	return 0, nil
}