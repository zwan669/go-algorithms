package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	ptr := head
	next := ptr.Next
	if ptr.Val == val {
		return next
	}
	for next != nil {
		if next.Val == val {
			ptr.Next = next.Next
		}
		ptr = next
		next = next.Next
	}
	return head
}
