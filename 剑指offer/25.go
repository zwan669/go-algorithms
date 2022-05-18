package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 先创建一个伪头节点
// l1 l2 元素依次比大小，小的连接到新链表的后面
// 最后把没有分配完的链表直接添加到新链表后面，然后返回伪头节点的下一个节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	ptr := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}
	if l1 == nil {
		ptr.Next = l2
	} else {
		ptr.Next = l1
	}
	return head.Next
}
