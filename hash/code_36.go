package hash

// https://leetcode.com/problems/valid-sudoku/description/
func isValidSudoku(board [][]byte) bool {
	for i, row := range board {
		for j, val := range row {
			c := string(val)
			if c == "." {
				continue
			}

			if !validNum(i, j, board) {
				return false
			}
		}
	}
	return true
}

func validNum(i, j int, board [][]byte) bool {
	minI := (i / 3) * 3
	maxI := minI + 2

	minJ := (j / 3) * 3
	maxJ := minJ + 2

	count := 0
	for x := minI; x <= maxI; x++ {
		for y := minJ; y <= maxJ; y++ {
			if board[x][y] == board[i][j] {
				count++
			}
		}
	}

	for x := 0; x < 9; x++ {
		if board[x][j] == board[i][j] {
			count++
		}
	}

	for y := 0; y < 9; y++ {
		if board[i][y] == board[i][j] {
			count++
		}
	}
	return count == 3
}
