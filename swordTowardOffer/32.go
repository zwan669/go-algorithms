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
	fmt.Println(levelOrder(tree))
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
