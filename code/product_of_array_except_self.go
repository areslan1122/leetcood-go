package main

import "fmt"

func productExceptSelf(nums []int) []int {

	nums_1 := make([]int, len(nums))
	nums_2 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums_1[i] = nums[i]
		nums_2[i] = nums[i]
	}

	for i := 1; i < len(nums); i++ {
		nums_1[i] = nums_1[i] * nums_1[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		nums_2[i] = nums_2[i] * nums_2[i+1]
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[i] = nums_2[i+1]
			continue
		}
		if i == len(nums)-1 {
			nums[i] = nums_1[i-1]
			continue
		}
		nums[i] = nums_1[i-1] * nums_2[i+1]
	}

	return nums

}

func main() {

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))

}
