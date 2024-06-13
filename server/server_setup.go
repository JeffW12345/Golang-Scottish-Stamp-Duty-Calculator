package server

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/JeffW12345/Golang-Scottish-Stamp-Duty-Calculator/server/server_model"
	"github.com/gin-gonic/gin"
)

var taxBands server_model.TaxBands = server_model.TaxBands{}
var taxCalculator server_model.TaxCalculator = server_model.TaxCalculator{}

func ServerSetup() {
	// Import tax bands and pass them to TaxCalculator object.
	taxBands.JsonConfigFilePath = "resources/tax_band_configuration.json"
	taxBands.ImportAndProcessTaxBands()
	taxCalculator.AddTaxBands(taxBands.Bands)

	// Set up and activate router
	router := gin.Default()
	router.GET("/get-tax-for-property-value/:propertyValue", getTaxForPropertValuePassedIn)
	router.Run("localhost:8080")
}

func getTaxForPropertValuePassedIn(c *gin.Context) {
	propertyValue := c.Param("propertyValue")

	propertyValueAsFloat, err := strconv.ParseFloat(strings.TrimSpace(propertyValue), 64)

	// If input not a number
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Need to provide a number"})
		return
	}
	// If input a number but out of range
	if propertyValueAsFloat < 0 || propertyValueAsFloat >= 2_147_483_647 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Need to provide a number"})
		return
	}
	// If successful request
	taxDue, _ := taxCalculator.CalculateTaxDue(propertyValueAsFloat)
	taxDueObject := server_model.TaxResult{TaxDue: taxDue}
	c.IndentedJSON(http.StatusOK, taxDueObject)
}