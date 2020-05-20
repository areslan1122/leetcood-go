package main

import "fmt"

/*
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

func fourSumCount(A []int, B []int, C []int, D []int) int {

	tmp := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			tmp[a+b]++
		}
	}

	count := 0
	for _, c := range C {
		for _, d := range D {
			count += tmp[-(c + d)]
		}
	}

	return count
}

func main() {

	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
