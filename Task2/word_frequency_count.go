package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func WordFrequency(input string) map[string]int {
	input = strings.ToLower(input)

	re := regexp.MustCompile(`[[:punct:]]+`)
	input = re.ReplaceAllString(input, "")

	words := strings.Fields(input)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	fmt.Println("Enter a string to analyze the word frequency:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	result := WordFrequency(input)
	fmt.Println("Word frequencies:", result)
}
