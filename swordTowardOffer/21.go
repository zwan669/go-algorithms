package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
	fmt.Println(exchange2([]int{1, 2, 3, 4}))
}

func exchange(nums []int) []int {
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	return nums
}

func exchange2(nums []int) []int {
	// 两个指针分别指向开头和末尾
	i, j := 0, len(nums)-1
	//左找偶右找奇，相遇就返回
	for i < j {
		for nums[i]%2 == 1 && i < j {
			i++
		}
		for nums[j]%2 == 0 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
