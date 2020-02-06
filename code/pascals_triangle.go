package main

import "fmt"

func generate(numRows int) [][]int {

	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{
			[]int{1},
		}
	}

	result := [][]int{[]int{1}}
	for i := 2; i <= numRows; i++ {
		newRow := make([]int, i)
		newRow[0] = 1
		newRow[i-1] = 1
		for j := 1; j < i-1; j++ {
			newRow[j] = result[len(result)-1][j-1] + result[len(result)-1][j]
		}
		result = append(result, newRow)
	}

	return result
}

func main() {

	fmt.Println(generate(6))
}
