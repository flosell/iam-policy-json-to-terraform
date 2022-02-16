let inputTextBox = document.getElementById("input")
let outputTextBox = document.getElementById("output")

let errorTextBox = document.getElementById("json-error")

let infoToggleButton = document.getElementById("info-toggle");
let infoExpander = document.getElementById("info-expander");

function convertToHcl() {

    let output
    try {
        output = convert("hello", inputTextBox.value)
        errorTextBox.textContent = ""
        errorTextBox.style.display = null
    } catch (e) {
        output = "Error: " + e.message
        errorTextBox.textContent = output
        errorTextBox.style.display = "block"
        return
    }
    errorTextBox.value = ""

    outputTextBox.textContent = output
    Prism.highlightElement(outputTextBox)

    window.goatcounter.count({
        path: 'convert-button-clicked',
        title: 'Convert-button was clicked',
        event: true,
    })

}
inputTextBox.addEventListener("input", convertToHcl)
let h = document.location.hash
let r = /content=(.+)/
let m = r.exec(h)
if (m) {
    inputTextBox.value = decodeURIComponent(m[1])

}
convertToHcl()
infoToggleButton.addEventListener("click",() => {
    infoExpander.toggleAttribute("open")
})