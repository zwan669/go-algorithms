package main

import "fmt"

func exist(board [][]byte, word string) bool {
	// result := false
	// i, j, k // i，j分别是board的行列索引，k是word的索引

	for i, v := range board {
		for j, _ := range v {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	fmt.Println("Search through, but no result found.")
	return false
}

func  dfs(board [][]byte, word string, i, j, k int) bool {
	/*
	判断i，j是否超出边界
	判断board[i][j]是否符合word[k]（判断是否访问过（通过把访问过的设置为空字符‘/000’））
	*/
	if i<0 || i>len(board)-1 || j<0 || j>len(board[0])-1 || board[i][j] != word[k] {
		return false
	}
	// 如果board[i][j]符合，word中的最后一个元素，返回true
	if k == len(word)-1 {
		return true
	}
	board[i][j] = '\000'
	// 上下左右依次检查, 搜索完毕记得把设置的'/000'改回来
	defer func() {board[i][j] = word[k]}()
	return dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i+1, j, k+1) ||
		dfs(board, word, i, j-1, k+1) ||
		dfs(board, word, i, j+1, k+1)
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))

	board = [][]byte{
		{'A', 'B'},
		{'C', 'D'},
	}
	word = "ABCD"
	fmt.Println(exist(board, word))
}