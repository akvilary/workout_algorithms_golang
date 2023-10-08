package main

import "fmt"

func main() {
	text := "()"
	result := isValid(text)
	fmt.Println(result)
}

func isValid(text string) bool {
	closeMapping := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	
	// fast check
	firstChar := string(text[0])
	if _, ok := closeMapping[firstChar]; ok {
		return false
	}

	// detailed check
	var stack []string
	var lastItem string
	for _, char := range text {
		closing, ok := closeMapping[string(char)]
		if !ok {
			stack = append(stack, string(char))
			continue
		}
		if len(stack) == 0 {
			return false
		}

		lastItem = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if closing != lastItem {
			return false
		}
	}

	return len(stack) == 0
}
