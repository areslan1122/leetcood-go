package main

//func merge(nums1 []int, m int, nums2 []int, n int) {
//
//	for i := 0; i < n; i++ {
//		nums1[m+i] = nums2[i]
//	}
//
//	for i := 0; i < m+n; i++ {
//		for j := i; j < m+n; j++ {
//			if nums1[i] > nums1[j] {
//				tmp := nums1[i]
//				nums1[i] = nums1[j]
//				nums1[j] = tmp
//			}
//		}
//	}
//
//}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}

	}

	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
