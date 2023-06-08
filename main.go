package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

var (
	lowercase_letters = "abcdefghijklmnopqrstuvwxyz";
	uppercase_letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	ambiguous_characters = "lLIO1";
	digits = "0123456789";
	special_characters = "!\"#$%&'*+,./:;=?@\\^`|~";
)

func main() {
	pass, err := generate_password(24, true, true, true, true);
	if err != nil {
		fmt.Println("Hello, Sailor!");
		return;
	}

	out := fmt.Sprintf("pass: %s", pass);
	fmt.Println(out);
}

func generate_password(length int, include_uppercase_letters bool, include_digits bool, include_special_characters bool, exclude_ambiguous_characters bool) (string, error) {
	if length < 1 {
		e := "password must be at least 1 character long";
		report_error(e)
		return "", errors.New(e);
	}

	// Get the available characters depending on params.

	available_characters := lowercase_letters;
	if include_digits { available_characters += uppercase_letters };
	if include_digits { available_characters += digits; }
	if include_special_characters { available_characters += special_characters; }

	// Remove the characters that may produce and "ambiguous" / hard-to-read passwords.
	if exclude_ambiguous_characters {
		for pos, char := range available_characters {
			if strings.ContainsRune(ambiguous_characters, char) {
				available_characters = removeCharAtIndex(available_characters, pos);
			}
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
	color_red := "\033[31m";
	color_reset := "\033[0m";
	fmt.Println(string(color_red) + "Keyed Error: " + error_message + "." + string(color_reset));
}

func removeCharAtIndex(str string, index int) string {
	if index < 0 || index >= len(str) {
		return str // Invalid index, return the original string
	}
	return str[:index] + str[index+1:]
}