package main

import "fmt"

func twoSum(num []int, target int) []int {
	save := map[int]int{}
	var output []int
	for index, value := range num {
		minus := target - value
		if index2, ok := save[minus]; !ok {
			save[value] = index
		} else {
			output = []int{index, index2}
		}
	}
	return output
}

func main() {
	num := []int{3, 2, 4}
	fmt.Println(twoSum(num, 6))
}
