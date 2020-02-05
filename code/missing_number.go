package main

import "fmt"

func missingNumber(nums []int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return (len(nums)*(len(nums)+1))/2 - sum
}

func main() {

	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4}))
}
