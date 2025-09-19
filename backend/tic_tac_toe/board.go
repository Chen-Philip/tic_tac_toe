package tic_tac_toe

import (
	"fmt"
)

type Board struct {
	board      [3][3]int
	winTracker [8]int
	isWin      bool
}

func (b *Board) resetBoard() {
	// Resets the board
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			b.board[i][j] = 0
		}
	}
	// Resets the tracker
	for i := 0; i < len(b.winTracker); i++ {
		b.winTracker[i] = 0
	}
	// Resets the win flag
	b.isWin = false
}

func (b *Board) makeMove(x, y int, isPlayer1 bool) {
	player := 1
	if !isPlayer1 {
		player = -1
	}

	// Updates the tracker
	isWin := b.updateTracker(x, player)
	isWin = isWin || b.updateTracker(3+y, player)
	if x+y == 2 {
		isWin = isWin || b.updateTracker(6, player)
	}
	if x == y {
		isWin = isWin || b.updateTracker(7, player)
	}

	// Updates and prints the board
	b.PrintBoard()
	b.board[x][y] = player
	b.isWin = isWin
}

func (b *Board) updateTracker(i, player int) bool {
	b.winTracker[i] += player
	return b.winTracker[i] == 3 || b.winTracker[i] == -3
}

func (b *Board) isValidMove(x, y int) bool {
	if x >= 3 || y >= 3 {
		fmt.Println("Input is out of bounds x y.")
		return false
	} else if x < 0 || y < 0 {
		fmt.Println("Input is out of bounds x y.")
		return false
	} else if b.board[x][y] != 0 {
		fmt.Println("This square is already used")
		return false
	}
	return true
}

func (b *Board) PrintBoard() {
	for i, row := range b.board {
		for j, cell := range row {
			move := " "
			if cell == -1 {
				move = "O"
			} else if cell == 1 {
				move = "X"
			}
			fmt.Print(" ", move, " ")
			// Prints the divider
			if j < len(row)-1 {
				fmt.Print("|")
			}
		}
		// Prints the divider
		if i < len(b.board)-1 {
			fmt.Println("\n---+---+---")
		}
	}
	fmt.Println()
}
