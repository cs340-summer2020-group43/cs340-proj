tableName = document.querySelector("table").getAttribute("data-db-table");

const deleteButtons = document.querySelectorAll('button.delete-row');

for (let deleteButton of deleteButtons) {
    deleteButton.onclick = (e) => {

        thisRow = deleteButton.closest("tr");

        testerId = thisRow.querySelector(".Tester").textContent
        productId = thisRow.querySelector(".Product").textContent

        e.preventDefault();

        let searchParams = new URLSearchParams();

        searchParams.append('table', tableName);
        searchParams.append('tester', testerId);
        searchParams.append('product', productId);

        fetch('/cs340/delete', {
            method: "POST",
            headers: {'Content-type': 'application/x-www-form-urlencoded'},
            body: searchParams,
        })
        .then((response) => {
            if (!response.ok) {
                throw Error(response.statusText);   
            }
            console.log(response);
            return response.text();
        })
        .then((result) => {
            console.log('success', result);
            thisRow.style.display = "none";
        })
        .catch((error) => {
            console.error('Error', error);
        })
    };
}
