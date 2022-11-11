package main

import (
	"fmt"
)

func main() {
	board := [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	activePlayer := "X"
	openFields := 9
	hasWinner := false

	for openFields > 0 && hasWinner == false {

		renderGameboard(board)
		var playersChoice int
		fmt.Scanln(&playersChoice)

		if checkLegalMove(playersChoice, board) {
			board = placeMarker(playersChoice, activePlayer, board)
			openFields--
			hasWinner = checkForWinner(board)
			activePlayer = switchPlayer(activePlayer)

		}
	}

	if hasWinner {

		fmt.Printf("%v hat gewonnen", switchPlayer(activePlayer))
	}

}

func renderGameboard(board [][]string) {

	fmt.Println("---+---+---")
	for i := 0; i < 3; i++ {
		fmt.Printf(" %v | %v | %v \n", board[i][0], board[i][1], board[i][2])
		fmt.Println("---+---+---")
	}

}

func checkLegalMove(field int, board [][]string) bool {

	/*
		switch field {
		case 1:
			return checkField(0, 0, board)
		case 2:
			return checkField(0, 1, board)
		case 3:
			return checkField(0, 2, board)
		case 4:
			return checkField(1, 0, board)
		case 5:
			return checkField(1, 1, board)
		case 6:
			return checkField(1, 2, board)
		case 7:
			return checkField(2, 0, board)
		case 8:
			return checkField(2, 1, board)
		case 9:
			return checkField(2, 2, board)

		default:
			return false
		}
	*/

	field--

	if field < 9 {

		return checkField(field/3, field%3, board)
	}
	return false
}

func checkField(x, y int, board [][]string) bool {
	return board[x][y] == " "

}

func placeMarker(field int, player string, board [][]string) [][]string {

	board[field/3][field%3] = player

	return board
}

func checkForWinner(board [][]string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != " " {
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != " " {
			return true
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[1][1] != " " {
		return true
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[1][1] != " " {
		return true
	}

	return false
}

func switchPlayer(activePlayer string) string {
	if activePlayer == "X" {
		return "O"
	}

	return "X"

}
