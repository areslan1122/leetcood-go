package main

import "fmt"
import "strings"

func reverseWords(s string) string {
	//	var str []string
	var ss string
	str := strings.Split(s, " ")
	for _, word := range str {
		ss += " "
		for i := len(word) - 1; i >= 0; i-- {
			ss += string(word[i])
		}
	}
	return ss[1:]
}

func main() {
	str := "Let's take LeetCode contest"
	fmt.Println(reverseWords(str))
}
