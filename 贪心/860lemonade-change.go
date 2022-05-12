package main

import "fmt"

func main() {
	bills := []int {5,5,5,10,20}
	fmt.Println(lemonadeChange(bills))
	bills = []int {5,5,10,10,20}
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	wallet := map[int]int{}
	for _, bill := range bills {
		switch bill {
		case 5:
			wallet[5]++
		case 10:
			if wallet[5] == 0 {
				return false
			}
			wallet[10]++
			wallet[5]--
		case 20:
			if wallet[10] != 0 && wallet[5] != 0 {
				wallet[20]++
				wallet[10]--
				wallet[5]--
			} else if wallet[5] >= 3 {
				wallet[20]++
				wallet[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}