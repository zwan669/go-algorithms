package main

import "fmt"

func main() {
	arr := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder(arr))
	arr = []int{1, 3, 2, 6, 5}
	fmt.Println(verifyPostorder(arr))
}

func verifyPostorder(postorder []int) bool {
	return helper(postorder, 0, len(postorder)-1)
}

func helper(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	mid := start                // 初始化mid，mid来找第一个比根节点大的元素
	root := postorder[end]      //根节点
	for postorder[mid] < root { // 找到第一个比根节点大的元素，找不到则mid等于root的下标
		mid++ // 此元素右边都是比根节点大
	} // 左边都比根节点小
	// ptr用于遍历右边数组，如果右边有比root小的则返回false
	for ptr := mid; ptr < end; ptr++ {
		if postorder[ptr] < root {
			return false
		}
	}
	return helper(postorder, start, mid-1) && helper(postorder, mid, end-1)
}
