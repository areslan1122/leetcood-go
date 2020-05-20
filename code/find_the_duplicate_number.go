package main

import "fmt"

/*
Input: [1,3,4,2,2]
Output: 2
*/
func findDuplicate(nums []int) int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}

	return 0
}

func main() {

	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2, 2, 2}))
}
