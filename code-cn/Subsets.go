package main

import "fmt"

func subsets(nums []int) [][]int {

	ans := make([][]int, 0)
	ans = append(ans, []int{})
	ansTmp := ans

	for _, v := range nums {
		for _, tmp := range ansTmp {

			ans = append(ans, append(append([]int{}, tmp...), v))
		}
		ansTmp = ans
	}
	return ans
}

func main() {

	nums := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(nums))

}
