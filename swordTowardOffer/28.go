package main

import (
	"fmt"
	. "github.com/zwan669/go-algorithms/swordTowardOffer/dataStructure"
)

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}
	fmt.Println(isSymmetric(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}
	// 思路双端队列
	if root.Left == nil || root.Right == nil {
		return false
	}
	deque := []*TreeNode{root.Left, root.Right}

	for len(deque) > 0 {
		left := deque[0]
		right := deque[len(deque)-1]
		deque = deque[1 : len(deque)-1]
		if left == right {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		// 左边插入
		// []*TreeNode{left.Right, left.Left} 相当于在队列头先添加左后添加右
		deque = append([]*TreeNode{left.Right, left.Left}, append(deque, right.Right, right.Left)...)
	}
	return true
}
func isSymmetric2(root *TreeNode) bool {
	var recur func(n1, n2 *TreeNode) bool
	recur = func(n1, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		return n1.Val == n2.Val && recur(n1.Left, n2.Right) && recur(n1.Right, n2.Left)
	}
	return recur(root, root)
}
