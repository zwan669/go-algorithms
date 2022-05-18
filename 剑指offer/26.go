package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var recur func(A, B *TreeNode) bool
	recur = func(A, B *TreeNode) bool {
		// 匹配结束条件：B的子节点为nil
		if B == nil {
			return true
		}
		// 不匹配结束条件：1.A为nil，B还有节点。
		// 			    2. A与B的值不匹配
		if A == nil || A.Val != B.Val {
			return false
		}
		return recur(A.Left, B.Left) && recur(A.Right, B.Right)
	}
	// DFS中序遍历：中左右
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// BFS
func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var recur func(A, B *TreeNode) bool
	recur = func(A, B *TreeNode) bool {
		// 匹配结束条件：B的子节点为nil
		if B == nil {
			return true
		}
		// 不匹配结束条件：1.A为nil，B还有节点。
		// 			    2. A与B的值不匹配
		if A == nil || A.Val != B.Val {
			return false
		}
		return recur(A.Left, B.Left) && recur(A.Right, B.Right)
	}
	var queue = []*TreeNode{A}
	for len(queue) > 0 {
		cur := queue[0]
		if recur(cur, B) {
			return true
		}
		queue = queue[1:]
		if cur == nil {
			continue
		}
		queue = append(queue, cur.Left, cur.Right)
	}
	return false
}
