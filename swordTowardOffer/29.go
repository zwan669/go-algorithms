package main

import "fmt"

func main() {
	var input = [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	res := spiralOrder(input)
	for i, _ := range res {
		fmt.Println(res[i])
	}
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var left, right, top, bottom int = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
	var res []int = make([]int, len(matrix)*len(matrix[0])) //创建 R*C 大小的数组
	var count int = 0
	for {
		// 左到右
		for i := left; i <= right; i++ {
			res[count] = matrix[top][i]
			count++
		}
		top++
		if top > bottom {
			break
		}
		// 上到下
		for i := top; i <= bottom; i++ {
			res[count] = matrix[i][right]
			count++
		}
		right--
		if left > right {
			break
		}
		// 右到左
		for i := right; i >= left; i-- {
			res[count] = matrix[bottom][i]
			count++
		}
		bottom--
		if top > bottom {
			break
		}
		// 下到上
		for i := bottom; i >= top; i-- {
			res[count] = matrix[i][left]
			count++
		}
		left++
		if left > right {
			break
		}
	}

	return res
}
