Interim development notes - to be deleted once the program has been completed.

MAIN THREAD
===========

A localhost server running on a separate thread will return the tax amount in response to a GET request (assuming that the response is a 2xx).

The main method would do the following: 

- Start the server in a goroutine.
- Use another goroutine to repeatedly ping the server with dummy requests until it responds with a 2xx status.
- Use a boolean channel to signal the main thread once the server is ready.

Once the channel reports that the server is ready, the main method then calls a void function of:

displayTaxDueForProperty(float32 valueOfProperty)

This function calls a function run of 'getTaxDueForPropertyOfValue(float32 valueOfProperty) float32, error'. 

This function sends off a GET request to obtain the value of the property. If returns:

- The value of the tax (or zero if the server was unable to process the request)

AND 

- Nil (or Error object if the server was unable to process the request

That information is used to provide a response to the user. 

SERVER ACTIONS
==============

The tax bands will be in a config file, stored in JSON:

[
    {"start": 0, "end": 145000, "percentageTax": 0},
    {"start": 145000.01, "end": 250000, "percentageTax": 0.02},
    {"start": 250000.01, "end": 325000, "percentageTax": 0.05},
    {"start": 325000.01, "end": 750000, "percentageTax": 0.10},
    {"start": 750000.01, "end": 2147483647, "percentageTax": 0.12}
]

When the app fires up, I import the band data from the config file. 

I create slice of Band objects.

I then sort the slice by start of band, ascending. 

I check the data integrity of the bands. Are the starts all lower than their ends? Is the start of this band 0.01 higher than the end of the last band? 

When a GET request is made, then if the GET doesn't request isn't for a positive number less than the max value of float32, then a 400 is returned. Otherwise, the tax is calculated as follows, in pseudocode:

taxDue = 0

for band in bands:
    if valueOfProperty <= band.start:
		break
	
	if valueOfProperty >= band.end:
		taxDue += ((band.start - band.end) * band.percentageTax)
	
	else:
		taxDue += ((valueOfProperty - band.start) * band.percentageTax)
	    break

The response body will contain a float representing the tax amount. I don't know yet if I'll need to wrap it in an object or if I can just return the number. 

FILES I NEED
============

tax_band_configuration.json (put in resources folder?)

'main' pacakge:

tax_band_calculator_entry_point.go - contains the main method. 

Following files to go into a package called 'server':

server_setup.go - Calls the functionality to import the tax bands and creates objects for them, and then creates a server that checks on a loop for new messages and acts on those messages. 
import_tax_bands.go - Contains a function with imports the tax bands from the config file and converts them into a slice of TaxBand objects. 
process_input.go - Contains a function to parse the JSON and check if it's valid (and return a 400 if not). Also contains a calculate_tax(valueOfProperty float32) float32 function. 

ASSUMPTIONS
===========

For the MVP, it will be assumed that:

- The tax amount will be entered into the main method. 
- This app will only ever be for Scottish property tax.
- The next tax band is a penny above the end of the previous tax band, not a pound as shown in https://revenue.scot/taxes/land-buildings-transaction-tax/residential-property#residential%20property%20rates%20and%20bands
- Rounding errors may occur (given that Go doesn't have a BigDecimal equivalent)


POSSIBLE AVENUES FOR FUTURE DEVELOPMENT
=======================================

- May get input from a user by a console or from localhost web page at some point.
- Web scraping https://revenue.scot/taxes/land-buildings-transaction-tax/residential-property#residential%20property%20rates%20and%20bands to check that our bands reflect what's shown there. 
- Adding functionality to cater for other taxes - maybe Stamp Duty for other parts of the UK.
- Program my own BigDecimal equivalent (in the unlikely event that there isn't one out there that I can use)