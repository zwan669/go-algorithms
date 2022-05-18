package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers2(2))
}

func printNumbers(n int) []int {
	nums := make([]int, int(math.Pow10(n)-1))
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}
	return nums
}

func printNumbers2(n int) string {
	chars := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var result string
	var number []byte = make([]byte, n)
	var myDfs func(x int)
	myDfs = func(x int) {
		if x == n {
			result += string(number)
			result += ","
			return
		}

		for _, char := range chars {
			number[x] = char
			myDfs(x + 1)
		}
	}
	myDfs(0)
	result = result[:len(result)-1]
	return result
}
