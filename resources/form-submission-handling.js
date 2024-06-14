document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('propertyValueForm').addEventListener('submit', async function(event) {
        event.preventDefault();

        const propertyValue = document.getElementById('valueInput').value;

        const url = '/get-tax-for-property-value/' + encodeURIComponent(propertyValue);

        try {
            const response = await fetch(url);
            if (!response.ok) {
                const errorMessageObject = await response.json();
                document.getElementById("results").style.display = "None";
                document.getElementById("error-handling").style.display = "block";
                document.getElementById("error-message").textContent = errorMessageObject.message
                throw new Error('Network response was not ok: ' + response.statusText);
            }

            const taxDueObject = await response.json();
            document.getElementById("valueOfProperty").textContent = String(parseFloat(propertyValue).toFixed(2)).replace(/\B(?=(\d{3})+(?!\d))/g, ",");
            document.getElementById("taxDue").textContent = String(parseFloat(taxDueObject.taxDue).toFixed(2)).replace(/\B(?=(\d{3})+(?!\d))/g, ",");
            document.getElementById("results").style.display = "block";
            document.getElementById("error-handling").style.display = "None";
        } catch (error) {
            console.error('Fetch error:', error);
        }
    });
});

//TODO - Code to handle unhappy path. 