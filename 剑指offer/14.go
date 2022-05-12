package main

import (
	"fmt"
)

func main() {
	fmt.Println(cuttingRope(56))
	fmt.Println(cuttingRope2(127))
	fmt.Println(cuttingRope3(127))
}

/*
时间复杂度：O(n^2)
空间复杂度：O(n)
转移方程：dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
*/
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1 // 初始值
	for i:=3; i<=n; i++ {
		for j:=2; j<i; j++ {
			max := 0
			if j*(i-j) > j*dp[i-j] {
				max = j*(i-j)
			} else {
				max = j*dp[i-j]
			}
			if dp[i] < max {
				dp[i] = max
			}
		}
	}
	return dp[n]
}

// 贪心

func cuttingRope2(n int) int {
	if n < 4 {
		return n-1
	}
	result := 1
	switch n%3 {
	case 0:
		for i:=0; i<n/3; i++ {
			result = result * 3 % (1e9+7)
		}
	case 1:
		for i:=0; i<n/3-1; i++ {
			result = result * 3 % (1e9+7)
		}
		result = result * 2 * 2 % (1e9+7)
	case 2:
		for i:=0; i<n/3; i++ {
			result = result * 3 % (1e9+7)
		}
		result = result * 2 % (1e9+7)
	}
	return result
}
// 快速幂求余
func cuttingRope3(n int) int {
	if n < 4 {
		return n-1
	}
	x, result := 3, 1
	// 乘积 3^(n//3） 减一是因为少成一位来判断余1和余2的情况
	for toThePowerOf := n/3-1 ;toThePowerOf != 0;toThePowerOf>>=1{
		if toThePowerOf & 1 == 1 {
			result = result * x % (1e9+7)
		}
		x = x * x % (1e9+7)
	}
	if remainder := n % 3; remainder == 0 {
		return result * 3 % (1e9+7)
	} else if remainder == 1 {
		return result * 4 % (1e9+7)
	} else {
		return result * 6 % (1e9+7)
	}
}