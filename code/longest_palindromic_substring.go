package main

import "fmt"

func longestPalindrome(s string) string {

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if checkPalindrome(s[j : len(s)-i+j]) {
				return s[j : len(s)-i+j]
			}
		}
	}
	return ""
}

func checkPalindrome(s string) bool {

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(longestPalindrome("aadadaaab"))
}
