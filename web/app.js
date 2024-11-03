let inputTextBox = document.getElementById("input")
let outputTextBox = document.getElementById("output")

let errorTextBox = document.getElementById("json-error")

let infoToggleButton = document.getElementById("info-toggle");
let infoExpander = document.getElementById("info-expander");

function convertToHcl() {

    let output
    try {
        output = window.ConvertString("hello", inputTextBox.value)
        errorTextBox.textContent = ""
        errorTextBox.style.display = null
    } catch (e) {
        output = "Error: " + e.message
        errorTextBox.textContent = output
        errorTextBox.style.display = "block"
        if (e.message && e.message.includes("looks like CloudFormation code")) {
            window.goatcounter.count({
                path: 'error-trying-to-convert-cloudformation',
                title: 'User is trying to convert CloudFormation code and its not supported',
                event: true,
            })
        }

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

function initialize() {
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
}

const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = 'wasm.wasm';

var wasm;

if ('instantiateStreaming' in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);
        initialize()
    })
} else {
    fetch(WASM_URL).then(resp =>
        resp.arrayBuffer()
    ).then(bytes =>
        WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
            wasm = obj.instance;
            go.run(wasm);
            initialize()
        })
    )
}
