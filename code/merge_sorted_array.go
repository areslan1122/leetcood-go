package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	for i := 0; i < m+n; i++ {
		for j := i; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				tmp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = tmp
			}
		}
	}

}
