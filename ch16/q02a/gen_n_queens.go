package ch16q02a

func GenNQueens(n int) [][]int {
	var result [][]int
	var rowArr []int
	solveNQueens(0, n, &rowArr, &result)
	return result
}

func solveNQueens(row, n int, rowArr *[]int, result *[][]int) {
	if row == n {
		// In Golang, make sure you make a copy of rowArr. Otherwise, the value of
		// of rowArr within result will keep getting updated!!!
		rowArrCopy := make([]int, len(*rowArr))
		copy(rowArrCopy, *rowArr)
		*result = append(*result, rowArrCopy)
	} else {
		for col := 0; col < n; col++ {
			*rowArr = append(*rowArr, col)
			if isValid(*rowArr) == true {
				solveNQueens(row+1, n, rowArr, result)
			}
			*rowArr = (*rowArr)[:len(*rowArr)-1]
		}
	}
}

func isValid(rowArr []int) bool {
	row := len(rowArr) - 1
	for i := 0; i < row; i++ {
		if rowArr[i] == rowArr[row] { // Check on same column
			return false
		}
		if abs(rowArr[i]-rowArr[row]) == row-i { // Check on same diagonal
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
