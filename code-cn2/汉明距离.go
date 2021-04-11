package main

import "fmt"

func hammingDistance(x int, y int) int {

	xor :=  x^y
	res := 0

	for {
		if xor == 0 {
			break
		}
		res += (xor % 2)
		xor /= 2
	}

	return res
}

func main() {
	fmt.Print(hammingDistance(4,8))
}