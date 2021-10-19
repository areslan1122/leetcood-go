package main

import "fmt"

func jump(nums []int) int {

	long := len(nums)
	dp := make([]int, long)

	for i:=0; i<long; i++ {
		for j := i+1; j<=i+nums[i] && j<long ; j++ {
			if dp[j] == 0 {
				dp[j] = dp[i] +1
			}
		}
	}

	return dp[long -1]
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
}
