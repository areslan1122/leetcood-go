package main

import "fmt"

/*
Input: [2,1,5,6,2,3]
Output: 10
*/

func largestRectangleArea(heights []int) int {

	max := 0
	tmp := 0
	for i := 0; i < len(heights); i++ {
		tmp = 1
		for j := i + 1; j < len(heights); j++ {
			if heights[j] >= heights[i] {
				tmp++
			} else {
				break
			}
		}

		for z := i - 1; z >= 0; z-- {
			if heights[z] >= heights[i] {
				tmp++
			} else {
				break
			}
		}

		if max < tmp*heights[i] {
			max = tmp * heights[i]
		}

	}

	return max

}

func main() {

	fmt.Println(largestRectangleArea([]int{2, 0, 2}))
}
