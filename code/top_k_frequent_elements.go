package main

import "fmt"

func topKFrequent(nums []int, k int) []int {

	res := []int{}
	Map := map[int]int{}

	for _, e := range nums {
		Map[e]++
	}

	for k > 0 {

		ress := 0
		index := -1

		for i, e := range Map {
			if e > index {
				index = e
				ress = i
			}

		}
		delete(Map, ress)
		res = append(res, ress)
		k--
	}

	return res
}

func main() {

	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

}
