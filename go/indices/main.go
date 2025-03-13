package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	indices := []int{}
	var sums int
	for k := 0; k < (len(nums) - 1); k++ {
		for n := 1; n < (len(nums) - k); n++ {

			sums = nums[k] + nums[k+n]
			if sums == target {
				indices = append(indices, k)
				indices = append(indices, k+n)
			}
		}
	}

	return indices
}

func main() {
	target := 12
	s := []int{3, 4, 5, 8}
	a := twoSum(s, target)
	fmt.Println(a)
}
