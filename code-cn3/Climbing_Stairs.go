package main

import "fmt"

//func climbStairs(n int) int {
//
//	if n == 1 {
//		return 1
//	} else if n == 2 {
//		return 2
//	}
//
//	return climbStairs(n-1) + climbStairs(n-2)
//}

func climbStairs(n int) int {

	tmp1 := 1
	tmp2 := 2
	ans := 0

	if n == 1 {
		return 1
	} else if n ==2 {
		return 2
	}

	for i := 3; i<= n; i++ {
		ans = tmp1 + tmp2
		tmp1 = tmp2
		tmp2 = ans
	}

	return ans
}


func main() {
	fmt.Println(climbStairs(44))
}