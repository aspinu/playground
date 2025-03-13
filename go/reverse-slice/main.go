package main

import (
	"fmt"
	"slices"
)

func rotate(num []int, k int) []int {
	fmt.Println(num[0])
	k1 := len(num) - k
	newNum := slices.Concat(num[k1:], num[:k1])
	return newNum
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	r := rotate(array, 3)
	fmt.Println(r)
}
