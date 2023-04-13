package main

import "fmt"

func shipWithinDays(weights []int, day int) int {

	l := 1
	r := 500 * 5 * 10000
	res := -1
	for l <= r {
		mid := (l + r) / 2
		countDay := 1
		curSum := 0
		for i := 0; i < len(weights); i++ {
			if curSum <= mid-weights[i] {
				curSum += weights[i]
			} else if weights[i] <= mid {
				curSum = weights[i]
				countDay++
			} else {
				countDay = 99999999999999
			}
		}
		if countDay <= day {
			r = mid - 1
			res = mid
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {
	//res := shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}
