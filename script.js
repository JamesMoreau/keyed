function updateNumberValue(value) {
	document.getElementById('lengthNumberInput').value = value;
}

function updateRangeValue(value) {
	document.getElementById('lengthRangeInput').value = value;
}

function copyPassword() {
	passwordInput = document.getElementById('PASS')

	password = passwordInput.value
	navigator.clipboard.writeText(password)
}

// WASM stuff

const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch('main.wasm'), goWasm.importObject)
.then((result) => {
	goWasm.run(result.instance)
})

function generatePasswordFromInput() {
	length = document.getElementById('lengthNumberInput').value
	includeUppercaseLettersInput = document.getElementById('CB0').checked
	includeDigitsInput = document.getElementById('CB1').checked
	includeSpecialCharactersInput = document.getElementById('CB2').checked
	excludeAmbiguousCharactersInput = document.getElementById('CB3').checked

	passwordInput = document.getElementById('PASS')
	password = generatePassword(length, includeUppercaseLettersInput, includeDigitsInput, includeSpecialCharactersInput, excludeAmbiguousCharactersInput)
	passwordInput.value = password
}
