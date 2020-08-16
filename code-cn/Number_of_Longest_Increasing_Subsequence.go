package main

import "fmt"

/*
Input: [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
*/

func findNumberOfLIS(nums []int) int {

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
	fmt.Println(dp)
	return max
}

func main() {

	arr := []int{1, 3, 5, 4, 7}
	fmt.Println(findNumberOfLIS(arr))
}
