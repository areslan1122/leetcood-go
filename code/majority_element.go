package main

import "fmt"

func majorityElement(nums []int) int {

	var a []int
	for i := 0; i < len(nums); i++ {

		if len(a) == 0 {
			a = append(a, nums[i])
		} else if a[len(a)-1] != nums[i] {
			a = a[:len(a)-1]
		} else {
			a = append(a, nums[i])
		}
	}
	return a[0]
}

func main() {

	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(arr))
}
