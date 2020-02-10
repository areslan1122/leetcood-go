package main

import "fmt"

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	max := 0
	if nums[1] < nums[0]+nums[2] {
		if len(nums) == 3 {
			return nums[0] + nums[2]
		} else {
			nums[2] = nums[0] + nums[2]
			max = nums[2]
		}

	} else {
		if len(nums) == 3 {
			return nums[1]
		} else {
			nums[2] = nums[0] + nums[2]
			max = nums[1]
		}
	}

	for i := 3; i < len(nums); i++ {

		if nums[i-2] > nums[i-3] {
			nums[i] = nums[i] + nums[i-2]
		} else {
			nums[i] = nums[i] + nums[i-3]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max

}

func main() {

	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
