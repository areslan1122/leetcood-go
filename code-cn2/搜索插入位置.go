package main

import "fmt"

func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) -1
	point := left + (right-left)/2

	for {
		fmt.Println(point,left,right,)
		if nums[point] == target {
			return point
		}

		if right - left <= 1 {
			if target <= nums[left] {
				return left
			}else if target <= nums[right] {
				return right
			} else if target > nums[right] {
				return right + 1
			}
		}

		if nums[point] < target {

			left = point
			point = left + (right -left)/2
		} else if nums[point] > target {
			right = point
			point = left + (right -left)/2
		}

	}

}

func main(){
	nums := []int{1,3}
	fmt.Println(searchInsert(nums, 3))
}