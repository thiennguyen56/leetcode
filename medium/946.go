package main

import "fmt"

func solution(pushed []int, popped []int) bool {
	j := 0
	stack := []int{}
	for _, val := range pushed {
		stack = append(stack, val)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return false
	}
	stack := []int{}
	i := 0
	j := 0
	for i < len(pushed) {
		fmt.Println(pushed[i], popped[j], stack)
		if len(stack) > 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		} else {
			stack = append(stack, pushed[i])
			i++
		}
	}
	for j < len(popped) {
		fmt.Println("2: ", stack, popped[j])
		if len(stack) > 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		} else {
			break
		}
	}
	return i == j
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(solution([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(solution([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
