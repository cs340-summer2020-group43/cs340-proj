const filterByFailButton = document.querySelector("button.filterByFail");

const filterHelperTextSpan = document.querySelector("span.filterHelperText");

const helperTextStrings = [
    "Showing All Tests",
    "Showing Only Failed Tests",
];

const buttonTextStrings = [
    "Filter out Passes",
    "Show all Tests",
];

let toggle = 0;

filterByFailButton.onclick = _ => {
    let rows = document.querySelectorAll('tbody tr');
    for (let row of rows) {
        if (row.querySelector(".pass").textContent == "PASS") {
            if (row.style.display == "none") {
                row.style.display = "";
            } else {
                row.style.display = "none";
            }
        }
    }

    toggle = (toggle + 1) % 2;
    filterHelperTextSpan.textContent = helperTextStrings[toggle];
    filterByFailButton.textContent = buttonTextStrings[toggle];
};

