package main

func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}
	count := 0
	for i := 1; i < len(nums); i++ {

		if nums[count] != nums[i] {
			nums[count+1] = nums[i]
			count++
		}
	}

	return count + 1
}
