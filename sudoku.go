package main

import "fmt"

const Size = 9

var board = [Size][Size]int{
	{0, 9, 6, 0, 4, 0, 0, 0, 1},
	{1, 0, 0, 0, 6, 0, 0, 0, 4},
	{5, 0, 4, 8, 1, 0, 3, 9, 0},
	{0, 0, 7, 9, 5, 0, 0, 4, 3},
	{0, 3, 0, 0, 8, 0, 0, 0, 0},
	{4, 0, 5, 0, 2, 3, 0, 1, 8},
	{0, 1, 0, 6, 3, 0, 0, 5, 9},
	{0, 5, 9, 0, 7, 0, 8, 3, 0},
	{0, 0, 3, 5, 9, 0, 0, 0, 7},
}

func findEmpty(board *[Size][Size]int) (int, int, bool) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}
func isValid(board *[Size][Size]int, row, col, num int) bool {
	for i := 0; i < Size; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func solve(board *[Size][Size]int) bool {
	row, col, empty := findEmpty(board)
	if !empty {
		return true // çözüm bulundu
	}

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solve(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}
func printBoard(board *[Size][Size]int) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}
func main() {
	if solve(&board) {
		fmt.Println("Sudoku Çözüldü:")
		printBoard(&board)
	} else {
		fmt.Println("Error.")
	}
}
