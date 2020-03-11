package main

func rotate(nums []int, k int) {

	for k1 := 0; k1 < k; k1++ {
		for i := len(nums) - 1; i > 0; i-- {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
