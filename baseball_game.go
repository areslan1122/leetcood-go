package main

import "fmt"
import "strconv"

func calPoints(ops []string) int {
	var nums []int
	sum := 0
	for _, i := range ops {
		if i == "C" {
			nums = nums[:len(nums)-1]
		} else if i == "D" {
			nums = append(nums, nums[len(nums)-1]*2)
		} else if i == "+" {
			nums = append(nums, nums[len(nums)-1]+nums[len(nums)-2])
		} else {
			num, _ := strconv.Atoi(i)
			nums = append(nums, num)
		}
	}
	for _, i := range nums {
		sum += i
	}

	return sum
}

func main() {
	var str []string
	str = append(str, "5", "-2", "4", "C", "D", "9", "+", "+")
	fmt.Println(calPoints(str))
}
