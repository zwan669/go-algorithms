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

// 双指针
// 第一遍loop记录链表长度
// 长度-k等于要走的步数， 例：长度=3 k=1 3-1=2那么就要走两步

func getKthFromEnd(head *ListNode, k int) *ListNode {
	length := 0
	ptr := head
	for ptr != nil {
		ptr = ptr.Next
		length++
	}

	for move := length - k; move > 0; move-- {
		head = head.Next
	}
	return head
}
