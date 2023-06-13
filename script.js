const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch('main.wasm'), goWasm.importObject)
	.then((result) => {
		goWasm.run(result.instance)
	})

function generatePasswordFromInput() {
	length = document.getElementById('LENGTH_NUMBER').value
	includeUppercaseLettersInput = document.getElementById('CB0').checked
	includeDigitsInput = document.getElementById('CB1').checked
	includeSpecialCharactersInput = document.getElementById('CB2').checked
	excludeAmbiguousCharactersInput = document.getElementById('CB3').checked

	// console.log(length, includeUppercaseLettersInput, includeDigitsInput, includeSpecialCharactersInput, excludeAmbiguousCharactersInput)

	passwordInput = document.getElementById('PASS')
	password = generatePassword(length, includeUppercaseLettersInput, includeDigitsInput, includeSpecialCharactersInput, excludeAmbiguousCharactersInput)
	passwordInput.value = password
}

function copyPassword() {
	passwordInput = document.getElementById('PASS')

	passwordInput.focus();
    passwordInput.select();
	
	password = passwordInput.value
	navigator.clipboard.writeText(password)
}