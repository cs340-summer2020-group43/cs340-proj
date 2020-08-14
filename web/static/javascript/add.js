const form = document.querySelector("form");
const tableName = form.getAttribute("data-db-table");

form.addEventListener('submit', (e) => {
    e.preventDefault();

    let formData = new FormData(form);
    let colValuePairs = new URLSearchParams();

    colValuePairs.append('table', tableName);
    for (var pair of formData) {
        if (pair[1] != "") {
            console.log(pair[0], pair[1]);
            colValuePairs.append(pair[0], pair[1]);
        }
    };

    console.log(colValuePairs);


    fetch('/cs340/insert', {
        method: "POST",
        headers: {
            'Content-type': 'application/x-www-form-urlencoded',
        },
        body: colValuePairs,
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
    })
    .catch((error) => {
        console.error('Error', error);
    })
});
