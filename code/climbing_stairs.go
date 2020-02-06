package main

import "fmt"

func climbStairs(n int) int {

	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	tmp := []int{1, 2}
	for i := 2; i < n; i++ {
		tmp = append(tmp, tmp[i-2]+tmp[i-1])
	}
	return tmp[n-1]

}

func main() {

	fmt.Println(climbStairs(5))
}
