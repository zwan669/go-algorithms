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

//
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr := head.Next
	head.Next = nil
	for {
		if ptr.Next == nil {
			ptr.Next = head
			return ptr
		}
		temp := ptr.Next
		ptr.Next = head
		head = ptr
		ptr = temp
	}
}

// 简洁版
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 递归版
func reverseList2(head *ListNode) *ListNode {
	var recur func(pre, cur *ListNode) *ListNode
	recur = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}
		temp := cur.Next
		cur.Next = pre
		return recur(cur, temp)
	}
	return recur(nil, head)
}
