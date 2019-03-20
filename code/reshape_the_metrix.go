package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	var arr1 []int
	var arr2 [][]int

	for _, i := range nums {
		arr1 = append(arr1, i...)
	}

	for i := 0; i < r; i++ {
		arr3 := arr1[0:c]
		arr1 = arr1[c:]
	}
	return arr2
}

func main() {

	a := [][]int{{1, 2}, {3, 4}}
	fmt.Println(a)
	fmt.Println(matrixReshape(a, 1, 4))
}
