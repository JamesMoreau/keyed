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
	lowercase_letters = "abcdefghijklmnopqrstuvwxyz";
	uppercase_letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	ambiguous_characters = "lLIO1";
	digits = "0123456789";
	special_characters = "!\"#$%&'*+,./:;=?@\\^`|~";
)

func generate_password(length int, include_uppercase_letters bool, include_digits bool, include_special_characters bool, exclude_ambiguous_characters bool) (string, error) {
	if length < 1 {
		e := "password must be at least 1 character long";
		report_error(e)
		return "", errors.New(e);
	}

	// Get the available characters depending on params.

	available_characters := lowercase_letters;
	
	if include_digits { 
		available_characters += uppercase_letters 
	}
	
	if include_digits { 
		available_characters += digits; 
	}

	if include_special_characters { 
		available_characters += special_characters; 
	}

	// Remove the characters that may produce and "ambiguous" / hard-to-read passwords. This must be done after adding all the characters together.
	if exclude_ambiguous_characters {
		for _, char := range ambiguous_characters {
			available_characters = strings.ReplaceAll(available_characters, string(char), "");
		}
	}

	var password_builder strings.Builder;
	available_characters_length := len(available_characters);

	for i := 0; i < length; i++ {
		big_int, err := rand.Int(rand.Reader, big.NewInt(int64(available_characters_length)))
		if err != nil {
			e := "unable to generate a random index";
			report_error(e);
			return "", errors.New(e);
		}

		char_index := big_int.Int64();

		character := available_characters[char_index];		
		err = password_builder.WriteByte(character);
		if err != nil {
			e := "unable to append rune";
			report_error(e);
			return "", errors.New(e);
		}
	}

	password := password_builder.String();
	return password, nil;
}

func report_error(error_message string) {
	fmt.Println("Keyed Error: " + error_message + ".");
}

func js_wrapper_generate_password(this js.Value, inputs []js.Value) interface{} {
	length_str := inputs[0].String();
	length, err := strconv.Atoi(length_str)
	if err != nil {
		report_error("Unable to convert length string to integer.");
	}

	include_uppercase_letters := inputs[1].Bool();
	include_digits := inputs[2].Bool();
	include_special_characters := inputs[3].Bool();
	exclude_ambiguous_characters := inputs[4].Bool();

	password, err := generate_password(length, include_uppercase_letters, include_digits, include_special_characters, exclude_ambiguous_characters);
	if err != nil {
		return "";
	}

	return password;
}

func main() {
	fmt.Println("Hello web assembly from go!");
	js.Global().Set("generatePassword", js.FuncOf(js_wrapper_generate_password));
	select {} // runs forever
}