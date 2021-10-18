package main

import "fmt"

//func uniquePaths(m int, n int) int {
//
//	if m ==1 && n == 1 {
//		return 1
//	}
//
//	if m == 1 {
//		return uniquePaths(m, n-1)
//	} else if n == 1 {
//		return uniquePaths(m-1, n)
//	}
//	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
//}

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)

	for i:=0; i<m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0 ; i< m ; i++ {
		for j := 0; j< n ; j++ {

			if i == 0 && j == 0 {
				dp[0][0] =1
			} else if i == 0 && j != 0 {
				dp[0][j] = dp[0][j -1]
			} else if i != 0 && j == 0 {
				dp[i][0] = dp[i-1][0]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}

	return dp[m-1][n-1]
}


func main() {
	fmt.Println(uniquePaths(7,3))
}