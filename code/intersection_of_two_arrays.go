package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {

	result := []int{}
	tmp := 0

	for _, v := range nums1 {
		if v == 0 {
			tmp++
			break
		}
	}
	for _, v := range nums2 {
		if v == 0 {
			tmp++
			break
		}
	}

	if tmp == 2 {
		result = append(result, 0)
	}

	for _, v := range nums1 {
		for _, v1 := range nums2 {
			if v == v1 && v != 0 {
				for k, v2 := range nums2 {
					if v == v2 {
						nums2[k] = 0
					}
				}
				result = append(result, v)
			}
		}
	}

	return result
}

func main() {

	a := []int{2, 2, 3, 4, 5, 6}
	b := []int{2, 2, 3, 4}
	fmt.Println(intersection(a, b))
}
