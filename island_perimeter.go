package main

import "fmt"

func islandPerimeter(grid [][]int) ([]int, [][]int, [][]int, int) {
	a := len(grid) + 2
	b := len(grid[0]) + 2
	grid_new := make([][]int, a)
	arr2 := grid_new
	arr := make([]int, b)

	for i := range arr {
		arr[i] = 0
	}
	for i := range grid_new {
		grid_new[i] = arr
	}
	arr2 = grid_new
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid_new[i+1][j+1] = grid[i][j]
		}
	}

	sum := 0

	for i := 1; i < len(grid_new)-1; i++ {
		for j := 1; j < len(grid_new[0])-1; j++ {
			if grid_new[i][j] == 1 {
				sum += 4
				sum -= grid_new[i-1][j]
				sum -= grid_new[i+1][j]
				sum -= grid_new[i][j+1]
				sum -= grid_new[i][j-1]
			}
		}
	}
	return arr, arr2, grid_new, sum

}

func main() {
	a := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(a))
}
