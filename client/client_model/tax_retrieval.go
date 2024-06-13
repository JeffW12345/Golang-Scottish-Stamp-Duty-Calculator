package client_model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"io"
)

// TaxRetrievalInterface is the interface for tax retrieval
type TaxRetrievalInterface interface {
	getTaxDue() (float32, error)
	displayTaxDueForProperty(valueOfProperty float32)
	isServerReadyYet() bool
}

type TaxRequest struct {
	PropertyValue float32 `json:"propertyValue"`
}

func (o *TaxRequest) getTaxDue() (float32, error) {
	response, err := http.Get(fmt.Sprintf("http://localhost:8080/get-tax-for-property-value/%v", o.PropertyValue))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		var er ErrorResponse
		if err := json.Unmarshal(bodyBytes, &er); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("Non-OK HTTP status: %v Details: %v", response.StatusCode, er.Message)
	}

	var tr TaxResponse
	if err := json.Unmarshal(bodyBytes, &tr); err != nil {
		return 0, err
	}

	return tr.TaxDue, nil
}

func (o *TaxRequest) DisplayTaxDueForProperty(valueOfProperty float32) {
	o.PropertyValue = valueOfProperty
	taxDue, err := o.getTaxDue()
	if err != nil {
		fmt.Println("Error retrieving tax:", err)
		return
	}
	fmt.Printf("The tax due for the property valued at %.2f is %.2f\n", valueOfProperty, taxDue)
}

func (o *TaxRequest) isServerReadyYet(startTime time.Time) bool {
	if time.Since(startTime) > 2*time.Second {
		fmt.Println("Server timed out")
		return false
	}
	_, err := o.getTaxDue()
	return err == nil
}

func (o *TaxRequest) WaitTillServerReady() {
	o.PropertyValue = 10
	startTime := time.Now()
	for {
		if o.isServerReadyYet(startTime) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
