package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	res := 0
	L, R := -1, -1
	min, max := math.MaxInt64, math.MinInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			R = i
			fmt.Println("R", i)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			L = i
			fmt.Println("L", i)
		}
	}
	if R > L {
		res = R - L + 1
	}
	return res
}

func main() {

	a := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
	fmt.Println("result", a)
}
