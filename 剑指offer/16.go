package main

import "fmt"

func main() {
	fmt.Println(myPow(2.10000, 3))

}

/*
-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104
*/
func myPow(x float64, n int) (ans float64) {
	if x == 0 {return 0}
	if n < 0 { // n为负数 则 x^n = （1/x）^(-n)
		x = 1 / x
		n = -n
	}
	ans = 1
	for; n!=0; n>>=1 {
		if n & 1 == 1 {
			ans *= x
		}
		x *= x // x^(2^0) => x^(2^0)*x^(2^0) = x^(2^1)
	}
	return
}