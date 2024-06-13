document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('propertyValueForm').addEventListener('submit', async function(event) {
        event.preventDefault();

        const propertyValue = document.getElementById('valueInput').value;

        const url = '/get-tax-for-property-value/' + encodeURIComponent(propertyValue);

        try {
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error('Network response was not ok: ' + response.statusText);
            }

            const taxDueObject = await response.json();
            document.getElementById("valueOfProperty").textContent = propertyValue;
            document.getElementById("taxDue").textContent = taxDueObject.taxDue;
            document.getElementById("results").style.display = "block";
        } catch (error) {
            console.error('Fetch error:', error);
        }
    });
});

//TODO - Code to handle unhappy path. 
// TODO - Code to format numbers with commas