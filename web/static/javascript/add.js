const myForm = document.getElementById('userAddForm');

myForm.addEventListener('submit', (e) => {
    e.preventDefault();
    //var userNameValue = document.getElementById('user_add_input_name').value;
    let formData = new FormData(myForm);
    let searchParam = new URLSearchParams();

    for (var pair of formData) {
        searchParam.append(pair[0], pair[1]);
    };

   fetch('https://postman-echo.com/post', {
        method: "POST",
        //headers: {'Content-type': 'application/x-www-form-urlencoded; charset=UTF-8'},
        cache: 'no-cache',
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
