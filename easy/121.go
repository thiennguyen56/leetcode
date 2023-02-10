package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	l := 0
	r := 1
	maxProfit := 0
	for l < len(prices) && r < len(prices) {
		profit := prices[r] - prices[l]
		if profit < 0 {
			l++
		} else {
			r++
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
func main() {
	fmt.Println("vim-go")
}
