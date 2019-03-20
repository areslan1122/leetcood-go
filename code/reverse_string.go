package main

import "fmt"

func reverseString(s string) string {
	var new_str string
	for i := len(s) - 1; i >= 0; i-- {
		new_str += string(s[i])
	}

	return new_str

}

func main() {
	a := "hello"
	fmt.Println(reverseString(a))
}
