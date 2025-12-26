let inputTextBox = document.getElementById("input")
let outputTextBox = document.getElementById("output")

let errorTextBox = document.getElementById("json-error")

let infoToggleButton = document.getElementById("info-toggle");
let infoExpander = document.getElementById("info-expander");

let copyInputButton = document.getElementById("copy-input");

function displayErrorMessage(errorMessage) {
    errorTextBox.textContent = "Error: " + errorMessage
    errorTextBox.style.display = "block"
}

function goatcounterCount(args) {
    if (window.goatcounter && window.goatcounter.count) {
        window.goatcounter.count(args)
    } else {
        console.log("goatcounter not loaded, maybe because of adblocker. Args: ", args)
    }
}

function countError(errorMessage) {
    if (errorMessage && errorMessage.includes("looks like CloudFormation code")) {
        goatcounterCount({
            path: 'error-trying-to-convert-cloudformation',
            title: 'User is trying to convert CloudFormation code and its not supported',
            event: true,
        })
    } else if (errorMessage && errorMessage.includes("did not contain any statements")) {
        goatcounterCount({
            path: 'error-lack-of-statements',
            title: 'User input didnt contain any statements',
            event: true,
        })
    } else if (errorMessage && errorMessage.includes("could not parse input")) {
        goatcounterCount({
            path: 'error-could-not-parse',
            title: 'User input couldnt be parsed as JSON',
            event: true,
        })
    } else {
        goatcounterCount({
            path: 'error-unknown',
            title: 'An error that we didnt expect happened',
            event: true,
        })
    }
}

function displayConversionResult(output) {
    setOutput(output)  

    errorTextBox.textContent = ""
    errorTextBox.style.display = null
    errorTextBox.value = ""
}

function convertToHcl() {
    try {
        let output = window.ConvertString("hello", inputTextBox.value)
        if (output.indexOf("Error") !== 0) {
            displayConversionResult(output);
        } else {
            displayErrorMessage(output);
            countError(output);
            return
        }

        goatcounterCount({
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

function copyToClipboard(text) {
    if (navigator.clipboard && window.isSecureContext) {
        return navigator.clipboard.writeText(text);
    } else {
        // Fallback for older browsers
        const textArea = document.createElement('textarea');
        textArea.value = text;
        textArea.style.position = 'fixed';
        textArea.style.left = '-999999px';
        textArea.style.top = '-999999px';
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();
        return new Promise((resolve, reject) => {
            document.execCommand('copy') ? resolve() : reject();
            textArea.remove();
        });
    }
}

function initialize() {
    inputTextBox.addEventListener("input", convertToHcl)
    let h = document.location.hash
    let r = /content=(.+)/
    let m = r.exec(h)
    if (m) {
        inputTextBox.value = decodeURIComponent(m[1])
        goatcounterCount({
            path: 'bookmarklet-used',
            title: 'Page was loaded using the bookmarklet',
            event: true,
        })

    }
    convertToHcl()
    infoToggleButton.addEventListener("click",() => {
        infoExpander.toggleAttribute("open")
        goatcounterCount({
            path: 'info-toggle-clicked',
            title: 'Info-Toggle-button was clicked',
            event: true,
        })
    })
    
    // Add copy functionality for input
    copyInputButton.addEventListener("click", async () => {
        try {
            await copyToClipboard(inputTextBox.value);
            const originalText = copyInputButton.textContent;
            copyInputButton.textContent = '✓';
            copyInputButton.style.background = '#4CAF50';
            setTimeout(() => {
                copyInputButton.textContent = originalText;
                copyInputButton.style.background = '';
            }, 1000);
            
            goatcounterCount({
                path: 'copy-input-clicked',
                title: 'Copy input button was clicked',
                event: true,
            });
        } catch (err) {
            console.error('Failed to copy text: ', err);
        }
    });
}

function setOutput(output) {
    const outputEl = document.getElementById("output")
    outputEl.innerHTML = ""
    
    // Create the code element
    const code = document.createElement("code")
    code.classList.add("language-hcl")  
    code.textContent = output
    
    // Create the pre wrapper that contains the code
    const preWrapper = document.createElement("pre")
    preWrapper.classList.add("language-hcl")
    preWrapper.appendChild(code)
    
    // Create copy button
    const copyButton = document.createElement("button")
    copyButton.textContent = "📋"
    copyButton.className = "copy-button"
    copyButton.title = "Copy HCL to clipboard"
    
    // Add copy functionality
    copyButton.addEventListener("click", async () => {
        try {
            await copyToClipboard(output);
            const originalText = copyButton.textContent;
            copyButton.textContent = '✓';
            copyButton.style.background = '#4CAF50';
            setTimeout(() => {
                copyButton.textContent = originalText;
                copyButton.style.background = '';
            }, 1000);
            
            goatcounterCount({
                path: 'copy-output-clicked',
                title: 'Copy output button was clicked',
                event: true,
            });
        } catch (err) {
            console.error('Failed to copy text: ', err);
        }
    });
    
    // Append pre to output container
    outputEl.appendChild(preWrapper)
    // Append button directly to outputEl 
    outputEl.appendChild(copyButton)
    
    // Highlight the code
    Prism.highlightElement(code)
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