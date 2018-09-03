package main

import "fmt"

func distributeCandies(candies []int) int {
	kind := make(map[int]int)
	for _, i := range candies {
		kind[i] = 1
	}
	if len(kind) >= len(candies)/2 {
		return len(candies) / 2
	} else {
		return len(kind)
	}
}

func main() {
	a := []int{1, 1, 2, 2, 3, 3, 3, 2}
	fmt.Println(distributeCandies(a))
}
