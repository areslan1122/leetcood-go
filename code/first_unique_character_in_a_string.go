package main

import "fmt"

func firstUniqChar(s string) int {

	if len(s) == 1 {
		return 0
	}

	str := []rune(s)

	for i := 0; i < len(str); i++ {
		tmp := true
		if str[i] == rune('-') {
			continue
		}
		for j := i; j < len(s); j++ {
			if str[i] == str[j] && (i != j) {
				str[j] = rune('-')
				tmp = false
			}
		}

		if !tmp {
			str[i] = rune('-')
			tmp = true
		}
	}

	for i := 0; i < len(str); i++ {
		if str[i] != rune('-') {
			return i
		}
	}

	return -1
}

func main() {

	fmt.Println(firstUniqChar("aabbabc"))
}
