const myForm = document.getElementById('errorAddButton');

myForm.addEventListener('submit', (e) => {
    e.preventDefault();
    let addForm = document.getElementById('errorAddForm');
    let formData = new FormData(addForm);
    let searchParam = new URLSearchParams();
    searchParam.append('table', 'Errors');
    console.log("check");
    
    for (var pair of formData) {
        searchParam.append(pair[0], pair[1]);
        console.log(pair[0], pair[1]);
    };

    fetch('/cs340/insert', {
        method: "POST",
        headers: {
            'Content-type': 'application/x-www-form-urlencoded',
            'Authorization': 'Client-ID [my-client-id]'
        },
        body: searchParam
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
