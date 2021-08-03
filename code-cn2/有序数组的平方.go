package main

import "fmt"

func sortedSquares(nums []int) []int {

	index_1 := 0
	index_2 := 0
	for i := 0; i< len(nums); i++ {
		if nums[i] >=0 {
			index_1 = i
			break
		}
		if i == len(nums) -1 {
			index_1 = len(nums)
		}
	}

	index_2 = index_1 -1
	res := []int{}
	for {
		if index_2 >= 0 && index_1 < len(nums) {

			res_1 := nums[index_1] * nums[index_1]
			res_2 := nums[index_2] * nums[index_2]
			if res_1 > res_2 {
				res = append(res, res_2)
				index_2--
			} else {
				res = append(res, res_1)
				index_1++
			}
		} else if index_2 < 0 && index_1 >= len(nums) {
			return res
		} else if index_2 < 0 {
			res = append(res, nums[index_1] * nums[index_1])
			index_1++
		} else if index_1 >= len(nums) {
			res = append(res, nums[index_2] * nums[index_2])
			index_2--
		}
	}


}

func main(){

	arr := []int{-4,-3,-2,-1}
	fmt.Println(sortedSquares(arr))
}
