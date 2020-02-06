package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	if len(cost) == 1 {
		return cost[0]
	} else if len(cost) == 2 && cost[0] < cost[1] {
		return cost[0]
	} else if len(cost) == 2 && cost[0] >= cost[1] {
		return cost[1]
	}

	for i := 2; i < len(cost); i++ {
		if cost[i-2] < cost[i-1] {
			cost[i] += cost[i-2]
		} else if cost[i-2] >= cost[i-1] {
			cost[i] += cost[i-1]
		}
	}

	if cost[len(cost)-1] > cost[len(cost)-2] {
		return cost[len(cost)-2]
	} else {
		return cost[len(cost)-1]
	}
}

func main() {

	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
