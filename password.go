package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"syscall/js"
)

var (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ambiguousCharacters = "lLIO1"
	digits = "0123456789"
	specialCharacters = "!\"#$%&'*+,./:=?@\\^`|~"
)

func generatePassword(length int, includeUppercaseLetters bool, includeDigits bool, includeSpecialCharacters bool, excludeAmbiguousCharacters bool) (string, error) {
	if length < 1 {
		e := "password must be at least 1 character long"
		reportError(e)
		return "", errors.New(e)
	}

	availableCharacters := lowercaseLetters
	
	if includeUppercaseLetters { 
		availableCharacters += uppercaseLetters 
	}
	
	if includeDigits { 
		availableCharacters += digits 
	}

	if includeSpecialCharacters { 
		availableCharacters += specialCharacters 
	}

	if excludeAmbiguousCharacters {
		for _, char := range ambiguousCharacters {
			availableCharacters = strings.ReplaceAll(availableCharacters, string(char), "")
		}
	}

	var passwordBuilder strings.Builder
	passwordBuilder.Grow(length)
	availableCharactersLength := len(availableCharacters)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(availableCharactersLength)))
		if err != nil {
			e := "unable to generate a random index"
			reportError(e)
			return "", errors.New(e)
		}

		charIndex := n.Int64()

		character := availableCharacters[charIndex]		
		err = passwordBuilder.WriteByte(character)
		if err != nil {
			e := "unable to append rune"
			reportError(e)
			return "", errors.New(e)
		}
	}

	return passwordBuilder.String(), nil
}

func reportError(error_message string) {
	fmt.Println("Keyed Error: " + error_message + ".")
}

func jsWrapperGeneratePassword(this js.Value, inputs []js.Value) interface{} {
	if len(inputs) < 5 {
		reportError("not enough arguments provided to jsWrapperGeneratePassword()")
		return ""
	}

	lengthStr := inputs[0].String()
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		reportError("Unable to convert length string to integer.")
	}

	includeUppercaseLetters := inputs[1].Bool()
	includeDigits := inputs[2].Bool()
	includeSpecialCharacters := inputs[3].Bool()
	excludeAmbiguousCharacters := inputs[4].Bool()

	password, err := generatePassword(length, includeUppercaseLetters, includeDigits, includeSpecialCharacters, excludeAmbiguousCharacters)
	if err != nil {
		return ""
	}

	return password
}

func main() {
	fmt.Println("Hello web assembly from go!")
	js.Global().Set("generatePassword", js.FuncOf(jsWrapperGeneratePassword))
	select {} // This runs forever so that the go program never exits.
}