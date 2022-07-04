package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 哈希表法
// 时间：O(N) 两轮遍历链表
// 空间：O(N) 哈希表 m 使用线性大小的额外空间。
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := map[*Node]*Node{}
	for ptr := head; ptr != nil; ptr = ptr.Next {
		m[ptr] = &Node{Val: ptr.Val}
	}
	for ptr := head; ptr != nil; ptr = ptr.Next {
		m[ptr].Next = m[ptr.Next]
		m[ptr].Random = m[ptr.Random]
	}
	return m[head]
}

//多了一次循环，但是空间复杂度变为O(1) 节点引用变量使用常数大小的额外空间。
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	//插入新节点
	for ptr := head; ptr != nil; ptr = ptr.Next.Next {
		node := &Node{Val: ptr.Val}
		ptr.Next, node.Next = node, ptr.Next
	}
	//处理random
	for ptr := head; ptr != nil; ptr = ptr.Next.Next {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		}
	}
	//分离两个链表
	copyHead := head.Next
	for ptr := head; ptr != nil; ptr = ptr.Next {
		node := ptr.Next
		ptr.Next = ptr.Next.Next
		if node.Next != nil {
			node.Next = node.Next.Next
		}
	}
	return copyHead
}
