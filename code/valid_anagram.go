package main

import "fmt"

// faster than 100%, memory 100%
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	axor := 0
	bxor := 0
	asum := 0
	bsum := 0

	for i := 0; i < len(s); i++ {
		axor = axor ^ int(s[i])
		asum = asum + (int(s[i]) * int(s[i]))
	}
	for i := 0; i < len(t); i++ {
		bxor = bxor ^ int(t[i])
		bsum = bsum + (int(t[i]) * int(t[i]))
	}

	if (axor != bxor) || (asum != bsum) {
		return false
	} else if axor^bxor == 0 {
		return true
	} else {
		return false
	}
}

// faster than 6.08%, memory less than 16.67%
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				if j == len(t)-1 {
					t = t[:j]
				} else {
					t = t[:j] + t[j+1:]
				}
				break
			}
		}
	}

	if t == "" {
		return true
	} else {
		return false
	}

}

func main() {

	s := "qwer"
	t := "rewq"
	fmt.Println(isAnagram(s, t))
}
