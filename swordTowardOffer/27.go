package main

import (
	. "github.com/zwan669/go-algorithms/swordTowardOffer/dataStructure"
)

func main() {

}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
