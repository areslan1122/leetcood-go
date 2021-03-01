package main

import "fmt"

func canPermutePalindrome(s string) bool {

	for i := 0 ; i<len(s) ; i++ {
		for j:=i+1 ; j<len(s); j++ {
			if s[i] == s[j] {
				s = s[:i] + s[i+1:]
				i--
				s = s[:j-1] + s[j:]
				break
			}
		}
	}

	return len(s) <= 1
}

func main() {

	fmt.Println(canPermutePalindrome("tactcoc"))
}