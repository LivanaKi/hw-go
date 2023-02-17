package main

import "fmt"

func main() {

	board := [3][3]int{
		{1, 1, 1},
		{0, 2, 2},
		{0, 0, 0}}
	fmt.Println(IsSolved(board))
}
func IsSolved(board [3][3]int) int {
	var res int
	if board[1][0] == board[1][1] && board[1][1] == board[1][2]{
		res = board[1][0]
	}
	if board[2][0] == board[2][1] && board[2][1] == board[2][2]{
		res = board[2][0]
	}
	if board[0][0] == board[1][0] && board[1][0] == board[2][0]{
		res = board[0][0]
	}
	if board[0][1] == board[1][1] && board[1][1] == board[2][1]{
		res = board[0][1]
	}
	if board[0][2] == board[1][2] && board[1][2] == board[2][2]{
		res = board[2][2]
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2]{
		res = board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0]{
		res = board[0][2]
	}
	if board[0][0] == board[0][1] && board[0][1] == board[0][2]{
		res = board[0][1]
	}
	for i := 0; i < len(board); i ++{
		for j := 0; j < len(board); j++{
			if res <= 0 && board[i][j] == 0{
				res = -1
			}
			if board[i][j] != 0 && res == 0{
				res = 0
			}
		}
	}

	return res
}
