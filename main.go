package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var (
	lowercase_letters = "abcdefghijklmnopqrstuvwxyz";
	uppercase_letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	digits = "0123456789";
	special_characters = "!\"#$%&'*+,./:;=?@\\^`|~";
)

func main() {
	fmt.Println("Hello, World");
	pass := generate_password(16, true, true);
	out := fmt.Sprintf("pass: %s", pass);
	fmt.Println(out);
}

func generate_password(length int, include_digits bool, include_special_characters bool) string {
	// get the avaiable characters depending on params.
	available_characters := lowercase_letters + uppercase_letters;
	if include_digits { available_characters += digits; }
	if include_special_characters { available_characters += special_characters; }

	var password_builder strings.Builder;
	available_characters_length := len(available_characters);

	for i := 0; i < length; i++ {
		big_int, err := rand.Int(rand.Reader, big.NewInt(int64(available_characters_length)))
		if err != nil {
			report_error("Unable to generate a random index.");
		}

		char_index := big_int.Int64();

		character := available_characters[char_index];		
		password_builder.WriteByte(character);
	}

	password := password_builder.String();
	return password;
}

func report_error(error_message string) {
	colorRed := "\033[31m"
	fmt.Println(string(colorRed), "Keyed Error: ", error_message);
}

