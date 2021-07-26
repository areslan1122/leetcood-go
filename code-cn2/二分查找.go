package main

import "fmt"

func search(nums []int, target int) int {


	index_1 := 0
	index_2 := len(nums)-1
	index := (index_1 + index_2)/2

	for {
		fmt.Println("one loop")
		if len(nums) == 0 {
			return -1
		}  else if  index_2 - index_1 == 1 || index_2 - index_1 == 0 {
			if nums[index_2] != target &&  nums[index_1] != target {
				return -1
			} else if nums[index_1] == target {
				return index_1
			} else if nums[index_2] == target {
				return index_2
			}
		}

		if target == nums[index] {
			return index
		}

		if target < nums[index] {
			index_2 = index
			index = (index_1 + index_2)/2
		} else {
			index_1 = index
			index = (index_1 + index_2)/2
		}
		fmt.Println(index, index_1, index_2)
	}
}

func main() {

	nums := []int{5}
	fmt.Println(search(nums,-5))
}
