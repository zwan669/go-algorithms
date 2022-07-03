package main

import (
	"fmt"
	. "github.com/zwan669/go-algorithms/swordTowardOffer/dataStructure"
)

func main() {
	tree := &TreeNode{
		1,
		&TreeNode{Val: 2},
		&TreeNode{
			3,
			&TreeNode{Val: 4},
			&TreeNode{Val: 5},
		},
	}
	fmt.Println(levelOrder3(tree))
}

func levelOrder(root *TreeNode) []int {
	sli, helperSli := []*TreeNode{root}, []*TreeNode{}
	var result []int
	for len(sli) > 0 {
		for i := 0; i < len(sli); i++ {
			cur := sli[i]
			result = append(result, cur.Val)
			if cur.Left != nil {
				helperSli = append(helperSli, cur.Left)
			}
			if cur.Right != nil {
				helperSli = append(helperSli, cur.Right)
			}
		}
		sli = helperSli
		helperSli = []*TreeNode{}
	}
	return result
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	sli, helperSli := []*TreeNode{root}, []*TreeNode{}
	var result [][]int
	var index int
	for len(sli) > 0 {
		result = append(result, make([]int, 0))
		for i := 0; i < len(sli); i++ {
			cur := sli[i]
			result[index] = append(result[index], cur.Val)
			if cur.Left != nil {
				helperSli = append(helperSli, cur.Left)
			}
			if cur.Right != nil {
				helperSli = append(helperSli, cur.Right)
			}
		}
		sli = helperSli
		helperSli = []*TreeNode{}
		index++
	}
	return result
}

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	sli, helperSli := []*TreeNode{root}, []*TreeNode{}
	var result [][]int
	var index int
	for len(sli) > 0 {
		result = append(result, make([]int, 0))
		for i := 0; i < len(sli); i++ {
			cur := sli[i]
			result[index] = append(result[index], cur.Val)
			if cur.Left != nil {
				helperSli = append(helperSli, cur.Left)
			}
			if cur.Right != nil {
				helperSli = append(helperSli, cur.Right)
			}
		}
		// reverse the array
		if index%2 == 1 {
			for i, j := 0, len(result[index])-1; i < j; i, j = i+1, j-1 {
				result[index][i], result[index][j] = result[index][j], result[index][i]
			}
		}
		sli = helperSli
		helperSli = []*TreeNode{}
		index++
	}
	return result
}
