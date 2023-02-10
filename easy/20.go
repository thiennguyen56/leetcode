package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	dict := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	for _, char := range s {
		opositeChar, ok := dict[char]
		if ok {
			stack = append(stack, opositeChar)
		} else if len(stack) > 0 {
			value := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if char != value {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println("vim-go")
}
