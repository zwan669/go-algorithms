package main

import (
	"fmt"
	. "github.com/zwan669/go-algorithms/swordTowardOffer/dataStructure"
)

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(pathSum(tree, 8))
}

func pathSum(root *TreeNode, target int) (res [][]int) {
	if root == nil {
		return
	}
	res = make([][]int, 0)
	var path []int
	var dfs func(node *TreeNode, tar, level int)
	dfs = func(node *TreeNode, tar, level int) {
		tar -= node.Val
		if len(path) <= level {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		if node.Left == nil && node.Right == nil && tar == 0 {
			res = append(res, append([]int(nil), path[:level+1]...))
			return
		}
		if node.Left != nil {
			dfs(node.Left, tar, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, tar, level+1)
		}
	}
	dfs(root, target, 0)
	return
}
