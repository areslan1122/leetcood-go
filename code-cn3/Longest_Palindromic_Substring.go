package main

import "fmt"

func longestPalindrome(s string) string {

	if s == "" {
		return ""
	}
	start, end := 0,0

	for i:=0; i< len(s); i++ {
		left1,right1 := helper(s, i, i)
		left2,right2 := helper(s, i, i+1)

		if right1 - left1 > end - start {
			start, end = left1, right1
		}

		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}

	return s[start:end+1]
}

func helper(s string, left, right int) (int,int) {

	for ; left >=0 && right <len(s) && s[left] == s[right]; left, right = left-1, right+1 {}
	return left+1, right-1
}

func main() {
	fmt.Println(longestPalindrome("ababcba"))
}
