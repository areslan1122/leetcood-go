package main

import "fmt"

func maxSubArray(nums []int) int {

	ans := nums[0]

	for i:= 1; i<len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}

		if nums[i] > ans {
			ans = nums[i]
		}
	}

	return ans
}

func main() {

	fmt.Println(maxSubArray([]int{1,2,-4,3,5,-2,5}))
}
