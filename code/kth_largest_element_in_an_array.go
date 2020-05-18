package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l < r {
		index := rand.Intn(r-l+1) + l
		nums[index], nums[r] = nums[r], nums[index]
		p, q := l, l
		for ; q < r; q++ {
			if nums[q] > nums[r] {
				nums[p], nums[q] = nums[q], nums[p]
				p++
			}
		}
		nums[p], nums[r] = nums[r], nums[p]
		// p numbers bigger than nums[p]
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			l = p + 1
		} else {
			r = p - 1
		}
	}
	return nums[k-1]
}

func main() {
	fmt.Println("vim-go")
}
