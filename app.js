let inputTextBox = document.getElementById("input")
let outputTextBox = document.getElementById("output")

let errorTextBox = document.getElementById("json-error")

let infoToggleButton = document.getElementById("info-toggle");
let infoExpander = document.getElementById("info-expander");

function displayErrorMessage(errorMessage) {
    errorTextBox.textContent = "Error: " + errorMessage
    errorTextBox.style.display = "block"
}

function countError(errorMessage) {
    if (errorMessage && errorMessage.includes("looks like CloudFormation code")) {
        window.goatcounter.count({
            path: 'error-trying-to-convert-cloudformation',
            title: 'User is trying to convert CloudFormation code and its not supported',
            event: true,
        })
    } else if (errorMessage && errorMessage.includes("did not contain any statements")) {
        window.goatcounter.count({
            path: 'error-lack-of-statements',
            title: 'User input didnt contain any statements',
            event: true,
        })
    } else if (errorMessage && errorMessage.includes("did not contain any statements")) {
        window.goatcounter.count({
            path: 'error-lack-of-statements',
            title: 'User input didnt contain any statements',
            event: true,
        })
    } else if (errorMessage && errorMessage.includes("could not parse input")) {
        window.goatcounter.count({
            path: 'error-could-not-parse',
            title: 'User input couldnt be parsed as JSON',
            event: true,
        })
    } else {
        window.goatcounter.count({
            path: 'error-unknown',
            title: 'An error that we didnt expect happened',
            event: true,
        })
    }
}

function displayConversionResult(output) {
    outputTextBox.textContent = output
    Prism.highlightElement(outputTextBox)

    errorTextBox.textContent = ""
    errorTextBox.style.display = null
    errorTextBox.value = ""
}

function convertToHcl() {
    try {
        let output = convert("hello", inputTextBox.value)

        displayConversionResult(output);

        window.goatcounter.count({
            path: 'convert-button-clicked',
            title: 'Convert-button was clicked',
            event: true,
        })
    } catch (e) {
        let errorMessage = e.message;
        displayErrorMessage(errorMessage);
        countError(errorMessage);
    }
}
inputTextBox.addEventListener("input", convertToHcl)
let h = document.location.hash
let r = /content=(.+)/
let m = r.exec(h)
if (m) {
    inputTextBox.value = decodeURIComponent(m[1])
    window.goatcounter.count({
        path: 'bookmarklet-used',
        title: 'Page was loaded using the bookmarklet',
        event: true,
    })

}
convertToHcl()
infoToggleButton.addEventListener("click",() => {
    infoExpander.toggleAttribute("open")
    window.goatcounter.count({
        path: 'info-toggle-clicked',
        title: 'Info-Toggle-button was clicked',
        event: true,
    })
})
