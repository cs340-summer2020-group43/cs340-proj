function validateForm() {
    var valid = true; // creates a boolean variable to return if the form's valid
    
    if (!validateField(this, 'userUpdateIdInput')) {// validates the name
        valid = false;
        document.getElementById("error").innerText = "Please enter an id value";
    }
    else {
        return myFunc();
    }
        return valid; // if all the fields are valid, this variable will be true
}

function validateField(context, fieldName) { // function to dynamically validates a field by its name
    var field = document.forms['userUpdateForm'][fieldName], // gets the field
        msg = 'Please enter the correct' + fieldName, // dynamic message
        errorField = document.getElementById(fieldName); // gets the error field
  console.log(context);
    // if the context is the form, it's because the Register Now button was clicked, if not, check the caller
    if (context instanceof HTMLFormElement || context.id === fieldName)
      errorField.innerHTML = (field.value === '') ? msg : '';
  
    return field.value !== ''; // return if the field is fulfilled
}

document.addEventListener('DOMContentLoaded', function() { // when the DOM is ready
    // add event handlers when changing the fields' value
    // document.getElementById('cellUpdateIdInput').addEventListener('input', validateForm);
    
    // add the event handler for the submit event
    document.getElementById('userUpdateButton').addEventListener('click', (e) => {
        e.preventDefault();
        validateForm();
        // document.getElementById('userUpdateIdInput').addEventListener('input', validateForm);
})});
 
function myFunc() {
    let myForm = document.getElementById("userUpdateForm");
    //e.preventDefault();
    
    let formData = new FormData(myForm);
    let searchParam = new URLSearchParams();

    searchParam.append('table', 'Users');

    for (var pair of formData) {
        searchParam.append(pair[0], pair[1]);
    };

    fetch('/cs340/update', {
        method: "POST",
        headers: {
            'Content-type': 'application/x-www-form-urlencoded',
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

};
// const myForm = document.getElementById('cellUpdateButton');

// function myForm.addEventListener('click', (e) => {
    // e.preventDefault();
// function myForm(e) {
//     e.preventDefault();
//     let addForm = document.getElementById('cellUpdateForm');
//     let formData = new FormData(addForm);
//     let searchParam = new URLSearchParams();
//     searchParam.append('table', 'Cells');
//     console.log("check");
    
//     for (var pair of formData) {
//         searchParam.append(pair[0], pair[1]);
//         console.log(pair[0], pair[1]);
//     };

//     fetch('/cs340/update', {
//         method: "POST",
//         headers: {
//             'Content-type': 'application/x-www-form-urlencoded',
//             'Authorization': 'Client-ID [my-client-id]'
//         },
//         body: searchParam
//         })
//         .then((response) => {
//             if (!response.ok) {
//                 throw Error(response.statusText);   
//             }
//             console.log(response);
//             return response.text();
//         })
//         .then((result) => {
//             console.log('success', result);
//         })
//         .catch((error) => {
//             console.error('Error', error);
//         })
// };
