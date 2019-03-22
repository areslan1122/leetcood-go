package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	ans := 1
	check := true
	arr := []string{}
	for i1 := 0; i1 < len(s)-1; i1++ {
		arr = append(arr, string(s[i1]))
		for i2 := i1 + 1; i2 < len(s); i2++ {
			for _, v := range arr {
				if string(s[i2]) == v {
					if len(arr) > ans {
						ans = len(arr)
					}
					check = false
					break
				}
			}
			if !check {
				check = true
				arr = []string{}
				break
			} else {
				arr = append(arr, string(s[i2]))
			}
			if i2 == len(s)-1 {
				if len(arr) > ans {
					ans = len(arr)
				}
				arr = []string{}
			}
		}
	}
	return ans
}

func main() {

	fmt.Println(lengthOfLongestSubstring(" "))

}
