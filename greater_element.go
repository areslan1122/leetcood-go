package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	for i := range nums {
		for j := 0; j < len(findNums); j++ {
			if nums[i] == findNums[j] {
				a := nums[i]
				nums[i] = -1
				for z := j; z < len(findNums); z++ {
					if findNums[z] > a {
						nums[i] = findNums[z]
						break
					}

				}
				break
			}
		}
	}
	return nums
}

func main() {
	a := []int{2, 4}
	b := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(b, a))
}
