package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(4294967293))
}

func hammingWeight(i uint32) int {
	i = (i & 0x55555555) + ((i >> 1) & 0x55555555)
	//步骤2
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	//步骤3
	i = (i & 0x0F0F0F0F) + ((i >> 4) & 0x0F0F0F0F)
	//步骤4
	i = i*(0x01010101) >> 24
	return int(i)
}

