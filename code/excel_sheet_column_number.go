package main

import "fmt"

func titleToNumber(s string) int {

	sum := 0
	for i := 0; i < len(s); i++ {

		sum = (sum * 26) + (int(s[i]) - 64)
	}

	return sum
}

func main() {

	s := "ABCD"
	titleToNumber(s)
	fmt.Println(titleToNumber(s))
}
