package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	l := 0
	r := len(nums) - 1
	res := -1
	for l <= r {
		mid := (l + r) / 2
		if (mid > 0 && nums[mid] != nums[mid-1]) && (mid < len(nums)-1 && nums[mid] != nums[mid+1]) {
			res = nums[mid]
			break
		} else if mid == 0 && nums[mid] != nums[mid+1] {
			res = nums[mid]
			break
		} else if mid == len(nums)-1 && nums[mid] != nums[mid-1] {
			res = nums[mid]
			break
		}
		if nums[mid] == nums[mid-1] {
			if mid+1 < len(nums) && len(nums[l:mid+1])%2 == 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			r = mid - 1
		}
	}
	return res
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3}))
}
