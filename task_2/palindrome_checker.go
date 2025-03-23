package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var cleanedString []rune
	for _, char := range s {
		if unicode.IsLetter(char) {
			cleanedString = append(cleanedString, char)
		}
	}

	for i := 0; i < len(cleanedString)/2; i++ {
		if cleanedString[i] != cleanedString[len(cleanedString)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Enter a string to check for palindrome:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	result := isPalindrome(input)

	if result {
		fmt.Printf("\"%s\" is a Palindrome.\n", input)
	} else {
		fmt.Printf("\"%s\" is NOT a Palindrome.\n", input)
	}
}
