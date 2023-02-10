package main

import "fmt"

func containsDuplicate(nums []int) bool {
	exist := map[int]bool{}
	for _, num := range nums {
		ex := exist[num]
		if !ex {
			exist[num] = true
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
