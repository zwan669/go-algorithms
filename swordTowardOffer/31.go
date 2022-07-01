package main

import "fmt"

func main() {
	var pushed, popped []int
	//pushed = []int{1, 2, 3, 4, 5}
	//popped = []int{4, 5, 3, 2, 1}
	//fmt.Println(validateStackSequences(pushed, popped))
	//pushed = []int{1, 2, 3, 4, 5}
	//popped = []int{4, 3, 5, 1, 2}
	//fmt.Println(validateStackSequences(pushed, popped))
	pushed = []int{1, 0}
	popped = []int{1, 0}
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	for i := 0; i < len(pushed); i++ {

		stack = append(stack, pushed[i])
		for j < len(popped) {
			pop := stack[len(stack)-1]
			if pop != popped[j] {
				break
			}
			j++
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
	}
	if j == len(popped) {
		return true
	}
	return false
}
