package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) == 0 {
		return getMedian(nums2)
	} else if len(nums2) == 0 {
		return getMedian(nums1)
	}



	ans := []int{}
	num1, num2 := 0,0

	for {

		if nums1[num1] > nums2[num2] {
			ans = append(ans, nums2[num2])
			num2++
			if len(nums2) == num2 {
				ans = append(ans, nums1[num1:]...)
				break
			}
		} else {
			ans = append(ans, nums1[num1])
			num1++
			if len(nums1) == num1 {
				ans = append(ans, nums2[num2:]...)
				break
			}
		}
	}
	return getMedian(ans)


}

func getMedian(ans []int)float64 {

	if len(ans) % 2 == 0 {
		return float64((ans[len(ans)/2] + ans[len(ans)/2-1])) / 2
	} else {
		return float64(ans[len(ans)/2])
	}
}


func main() {
	a := []int{1,2}
	b := []int{3,4}

	fmt.Println(findMedianSortedArrays(a,b))

}