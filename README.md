INTRODUCTION
============

This app calculates LBBT (essentially Scottish stamp duty - see https://revenue.scot/taxes/land-buildings-transaction-tax).

The user passes the property value into the DisplayTaxDueForProperty method in the main function of 'main.go', and is presented with 
console output showing the tax due on the property. In the £200,000 property example already in the app, the console output would be the following:

The tax due for the property valued at 200000.00 is 1100.00

RUNNING INSTRUCTIONS
====================

Clone the repo. 

Navigate to the repo in the command terminal. Then type the following:

go get .

Change the figure in the 'DisplayTaxDueForProperty' method call in the 'main' function to the value of the property you want the tax for. 

Enter the following:

go run .

ASSUMPTIONS
===========

For the MVP, it will be assumed that:

- This app will only ever be for Scottish property tax.
- Rounding errors may occur
- The value of the property is a positive number <= 2,147,483,647.
- The only consideration when calculating the tax is the property value. 


POSSIBLE AVENUES FOR FUTURE DEVELOPMENT
=======================================

- May get input from a user by a console or from localhost web page.
- Web scraping https://revenue.scot/taxes/land-buildings-transaction-tax/residential-property#residential%20property%20rates%20and%20bands to check that the app's bands reflect what's shown there. 
- Adding functionality to cater for other taxes - maybe Stamp Duty for other parts of the UK.
- Using a decimal library to avoid rounding errors. 