package main

import (
	"fmt"
)

type Player int

const (
	Empty   Player = iota // 0
	Player1               // 1
	Player2               // 2
)

type Board [6][7]Player

func main() {
	var board Board
	var currentPlayer Player = Player1

	for {
		fmt.Println("")
		fmt.Println("État actuel du plateau :")
		fmt.Println("")
		printBoard(board)

		column := askPlayerColumn(board, currentPlayer)
		row := getLastEmptyRow(board, column)

		board[row][column] = currentPlayer

		if gameIsOver(board, column, row) {
			fmt.Printf("\nLa partie est terminée. Le Joueur %d a gagné !\n\n", currentPlayer)
			printBoard(board)
			fmt.Println("")
			break
		}
		currentPlayer = switchPlayer(currentPlayer)

	}
}

func askPlayerColumn(b Board, currentPlayer Player) int {
	choice := -1
	for {
		fmt.Printf("\nC'est au tour du Joueur %d. Choisissez une colonne : ", currentPlayer)
		fmt.Scan(&choice)
		if choice >= 0 && choice <= 6 {
			if getLastEmptyRow(b, choice) != -1 {
				return choice
			}
		}
		fmt.Println("\nVeuillez choisir une colonne valide.")
		printBoard(b)
	}
}

func getLastEmptyRow(b Board, column int) int {
	for i := 0; i < 6; i++ {
		if b[i][column] != Player1 && b[i][column] != Player2 {
			return i
		}
	}
	return -1
}

func printBoard(b Board) {
	// Afficher le plateau ici
	fmt.Println(" | 0 | 1 | 2 | 3 | 4 | 5 | 6 |")
	fmt.Println("-------------------------------")
	for i := 5; i >= 0; i-- {
		fmt.Printf("%d|", i)
		for j := 0; j < 7; j++ {
			switch b[i][j] {
			case Player1:
				fmt.Printf(" X |")
			case Player2:
				fmt.Printf(" O |")
			default:
				fmt.Printf(" . |")
			}
		}
		fmt.Println("")
	}

}
func switchPlayer(current Player) Player {
	switch current {
	case Player1:
		return Player2
	case Player2:
		return Player1
	default:
		return Empty
	}

}
func gameIsOver(b Board, column int, row int) bool {
	if checkWinRow(b, column, row) || checkWinColumn(b, column, row) || checkWinDiagonal(b, column, row) {
		return true
	}
	return false
}

func checkWinRow(b Board, column int, row int) bool {
	if column <= 3 && b[row][column] == b[row][column+1] && b[row][column+1] == b[row][column+2] && b[row][column+2] == b[row][column+3] {
		return true
	}

	if column >= 3 && b[row][column] == b[row][column-1] && b[row][column-1] == b[row][column-2] && b[row][column-2] == b[row][column-3] {
		return true
	}
	return false
}

func checkWinColumn(b Board, column int, row int) bool {
	if row <= 2 && b[row][column] == b[row+1][column] && b[row+1][column] == b[row+2][column] && b[row+2][column] == b[row+3][column] {
		return true
	}
	if row >= 3 && b[row][column] == b[row-1][column] && b[row-1][column] == b[row-2][column] && b[row-2][column] == b[row-3][column] {
		return true
	}
	return false
}

func checkWinDiagonal(b Board, column int, row int) bool {
	if column <= 3 && row >= 3 && b[row][column] == b[row-1][column-1] && b[row-1][column-1] == b[row-2][column-2] && b[row-2][column-2] == b[row-3][column-3] {
		return true
	}

	if column >= 3 && row <= 2 && b[row][column] == b[row+1][column+1] && b[row+1][column+1] == b[row+2][column+2] && b[row+2][column+2] == b[row+3][column+3] {
		return true
	}
	return false
}
