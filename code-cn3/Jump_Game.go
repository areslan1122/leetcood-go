package main

func canJump(nums []int) bool {

	tmp := nums[0]

	for i:=0; i<len(nums); i++ {

		if tmp < i {
			break
		}
		if nums[i] + i > tmp {
			tmp = nums[i] + i
		}
	}

	if tmp >= len(nums) -1 {
		return true
	} else {
		return false
	}

}
