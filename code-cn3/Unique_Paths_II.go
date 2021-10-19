package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m,n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([]int, n)
		dp[0] = 1

	for i :=0; i < m; i++ {
		for j := 0; j < n; j++ {

			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				fmt.Println(i,j,dp[j])
				continue
			}

			if j -1 >= 0 && obstacleGrid[i][j] == 0 {
				dp[j] += dp[j-1]
			}
			fmt.Println(i,j,dp[j])
		}
	}

	return dp[n-1]
}

func main() {

	fmt.Println(uniquePathsWithObstacles([][]int{[]int{0,0,0},[]int{0,1,0},[]int{0,0,0}}))

}