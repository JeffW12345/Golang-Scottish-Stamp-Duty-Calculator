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

    // Serve the home.html file when visiting the root URL
    router.GET("/", func(c *gin.Context) {
        c.File("./resources/home.html")
    })

    // Serve static files (like CSS, JS) from the resources directory
    router.Static("/resources", "./resources")

    // Handle the tax calculation endpoint
    router.GET("/get-tax-for-property-value/:propertyValue", getTaxForPropertValuePassedIn)

    router.Run("localhost:8080")
}


func getTaxForPropertValuePassedIn(c *gin.Context) {
	propertyValue := c.Param("propertyValue")

	propertyValueAsFloat, err := strconv.ParseFloat(strings.TrimSpace(propertyValue), 64)

	// If input not a number
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "need to provide a number between 0 and 2,147,483,647"})
		return
	}
	// If input a number but out of range
	if propertyValueAsFloat < 0 || propertyValueAsFloat >= 2_147_483_647 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "need to provide a number between 0 and 2,147,483,647"})
		return
	}
	// If successful request
	taxDue, _ := taxCalculator.CalculateTaxDue(propertyValueAsFloat)
	taxDueObject := server_model.TaxResult{TaxDue: taxDue}
	c.IndentedJSON(http.StatusOK, taxDueObject)
}