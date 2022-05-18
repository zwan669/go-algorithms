package main

import "fmt"

func numDigit(num int) (res int) {
	for num != 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func movingCount(m int, n int, k int) (res int) {
	// res = bfs(m, n, k)
	res = DFS(m, n, k, 0, 0, make(map[[2]int]bool))
	return
}

func bfs(m, n, k int) int {
	// 创建queue把0，0的起始点加进去
	var queue = [][2]int {{0, 0}}
	// 创建一个储存访问过元素的set
	var visited = map[[2]int]bool{}
	for len(queue) != 0 {
		point := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// 如果当前行数大于m
		// 或列数大于n
		// 或行列位数之和大于k
		// 或访问过
		// 则继续pop队列中的元素（继续搜索）
		if point[0]>=m || point[1]>=n || numDigit(point[0]) + numDigit(point[1])>k || visited[point] {continue}
		visited[point] = true
		// 添加元素到队列最前方
		queue = append([][2]int{{point[0]+1, point[1]}, {point[0], point[1]+1}}, queue...)
	}
	return len(visited)
}

func DFS(m, n, k, row, col int, visited map[[2]int]bool) int {
	// 基本结束条件：
	// 行数大于m
	// 列数大于n
	// 行列位数之和大于k
	if row >= m || col >= n || numDigit(row) + numDigit(col) > k || visited[[2]int{row, col}] {
		return 0
	}
	visited[[2]int{row, col}] = true
	return DFS(m, n, k, row+1, col, visited) + DFS(m, n, k, row, col+1, visited) + 1
}
func main() {
	//fmt.Println(movingCount(2, 3, 1))
	//fmt.Println(movingCount(3, 1, 0))
	fmt.Println(movingCount(38, 15, 9))
}


