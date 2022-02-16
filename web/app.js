let inputTextBox = document.getElementById("input")
let outputTextBox = document.getElementById("output")
let errorTextBox = document.getElementById("error")

function convertToHcl() {
    let output
    try {
        output = convert("hello", inputTextBox.value)
    } catch (e) {
        output = "Could not convert:\n" + e.message
        errorTextBox.value = output

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