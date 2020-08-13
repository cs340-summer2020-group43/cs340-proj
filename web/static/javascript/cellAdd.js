const addForm = document.getElementById('cellAddForm');

addForm.addEventListener('submit', (e) => {
    e.preventDefault();

    let formData = new FormData(addForm);
    let searchParams = new URLSearchParams();

    searchParams.append('table', 'Cells');
    
    for (var pair of formData) {
        searchParams.append(pair[0], pair[1]);
    };

    fetch('/cs340/insert', {
        method: "POST",
        headers: {
            'Content-type': 'application/x-www-form-urlencoded',
        },
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
    })
    .catch((error) => {
        console.error('Error', error);
    })
});
