package main

import "fmt"

/*
Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))

	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {

	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(arr))
}
