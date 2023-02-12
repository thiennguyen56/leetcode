package main

import "fmt"

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	for i := 1; i < length; i++ {
		if nums[i] > maxEndingHere+nums[i] {
			maxEndingHere = nums[i]
		} else {
			maxEndingHere += nums[i]
		}
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func main() {
	fmt.Println("vim-go")
}
