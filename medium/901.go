package main

import "fmt"

type StockSpanner struct {
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, price)
		return 1
	}
	res := 0
	for i := len(this.stack) - 1; i >= 0; i-- {
		if price >= this.stack[i] {
			res += 1
		} else {
			break
		}
	}

	this.stack = append(this.stack, price)
	return res + 1
}

func main() {
	fmt.Println("vim-go")
}
