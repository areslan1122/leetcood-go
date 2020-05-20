package main

/*
Input:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

*/
func setZeroes(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] > 0 {
				matrix[i][j]++
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				for y := 0; y < len(matrix); y++ {
					if matrix[y][j] != 0 {
						matrix[y][j] = 1
					}
				}
				for y := 0; y < len(matrix[i]); y++ {
					if matrix[i][y] != 0 {
						matrix[i][y] = 1
					}
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] > 0 {
				matrix[i][j]--
			}
		}
	}

}
