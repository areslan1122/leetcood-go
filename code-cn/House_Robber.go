package main

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if i > 1 {
			nums[i] += nums[i-2]
		}
		if nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
	}

	return nums[len(nums)-1]
}
