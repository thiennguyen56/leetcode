package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS := map[rune]int{}
	countT := map[rune]int{}

	for i := 0; i < len(s); i++ {
		countS[rune(s[i])] += 1
		countT[rune(t[i])] += 1
	}

	return reflect.DeepEqual(countS, countT)
}

func main() {
	fmt.Println("vim-go")
}
