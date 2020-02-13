package main

import "fmt"

func isPalindrome(s string) bool {

	str := []byte{}
	for i := 0; i < len(s); i++ {
		if (s[i] < 91 && s[i] > 64) || (s[i] <= 57 && s[i] >= 48) {
			str = append(str, s[i])
		}
		if s[i] < 123 && s[i] > 96 {
			str = append(str, s[i]-32)
		}
	}
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(isPalindrome("0p"))
}
