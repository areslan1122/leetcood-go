package main

import "fmt"

func main() {

	a := 4
	b := 4
	grid_new := make([][]int, a)
	arr := make([]int, b)
	for i := range grid_new {
		grid_new[i] = arr
	}
	fmt.Println(arr)
	fmt.Println(grid_new)
}
