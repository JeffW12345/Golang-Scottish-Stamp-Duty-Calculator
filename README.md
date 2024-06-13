INTRODUCTION
============

This app calculates LBBT (essentially Scottish stamp duty - see https://revenue.scot/taxes/land-buildings-transaction-tax).

The calculations are done on a localhost server. The client obtains the tax amount for a property of a given value by making a GET request to 
the server's API.

RUNNING INSTRUCTIONS
====================

Clone the repo. 

Navigate to the repo in a command terminal. Then type the following:

go mod download

Now enter the following:

go run .

You will now be presented with console output which includes the following:

**CLICK HERE: http://localhost:8080/**

Press control and then left click on the hyperlink. If you enter the property value and press 'submit', you will be presented with the tax amount.

ASSUMPTIONS
===========

For the MVP, it will be assumed that:

- This app will only ever be for Scottish property tax.
- Rounding errors may occur
- The value of the property is a positive number <= 2,147,483,647.
- The only consideration when calculating the tax is the property value. 


POSSIBLE AVENUES FOR FUTURE DEVELOPMENT
=======================================
- Web scraping https://revenue.scot/taxes/land-buildings-transaction-tax/residential-property#residential%20property%20rates%20and%20bands to check that the app's bands reflect what's shown there. 
- Adding functionality to cater for other taxes - maybe Stamp Duty for other parts of the UK.
- Using a decimal library to avoid rounding errors. 
