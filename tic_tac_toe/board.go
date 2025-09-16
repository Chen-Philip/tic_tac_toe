package tic_tac_toe

import (
	"fmt"
)

type Board struct {
	board [3][3]int
	State BoardState
}

func (b Board) PrintBoard() {
	for i, row := range b.board {
		for j, cell := range row {
			fmt.Print(" ", cell, " ")
			if j < len(row)-1 {
				fmt.Print("|")
			}
		}
		if i < len(b.board)-1 {
			fmt.Println("\n---+---+---")
		}
	}
}
