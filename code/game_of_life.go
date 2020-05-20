package main

/*
0: d -> d
1: a -> a
2: a -> d
3: d -> a
*/

func gameOfLife(board [][]int) {

	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			alive := 0
			alive += isLive(board, i-1, j)
			alive += isLive(board, i+1, j)
			alive += isLive(board, i, j-1)
			alive += isLive(board, i, j+1)
			alive += isLive(board, i+1, j-1)
			alive += isLive(board, i+1, j+1)
			alive += isLive(board, i-1, j-1)
			alive += isLive(board, i-1, j+1)

			if board[i][j] == 0 {
				if alive == 3 {
					board[i][j] = 3
				}
			} else if board[i][j] == 1 {
				if alive > 3 || alive < 2 {
					board[i][j] = 2
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] %= 2
		}
	}
}

func isLive(board [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 0 || board[i][j] == 3 {
		return 0
	}
	return 1
}
