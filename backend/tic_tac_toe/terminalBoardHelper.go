package tic_tac_toe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (g *Game) getMoveFromTerminal() (int, int) {
	x, y := -1, -1
	for x == -1 && y == -1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Make your move (x y): ")
		line, _ := reader.ReadString('\n')
		coords := strings.Fields(line)
		if len(coords) != 2 {
			fmt.Println("Please make a valid move in the form of x y.")
		} else {
			num1, err1 := strconv.Atoi(coords[0])
			num2, err2 := strconv.Atoi(coords[1])

			if (err1 != nil) || (err2 != nil) {
				fmt.Println("The inputs have to be numbers.")
			} else if g.isOnBoard(num2, num1) {
				x = num1
				y = num2
			}
		}
	}
	return y, x
}

func (g *Game) isOnBoard(x, y int) bool {
	if x >= 3 || y >= 3 {
		fmt.Println("Input is out of bounds x y.")
		return false
	} else if x < 0 || y < 0 {
		fmt.Println("Input is out of bounds x y.")
		return false
	}
	return g.isValidMove(x, y)
}

func (g *Game) PrintBoard() {
	for i, row := range g.board {
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
		if i < len(g.board)-1 {
			fmt.Println("\n---+---+---")
		}
	}
	fmt.Println()
}

func (g *Game) StartTerminalGame() {
	hasWinner := false
	i := 0
	for !hasWinner && i < 9 {
		x, y := g.getMoveFromTerminal()
		g.makeMove(x, y, i%2 == 0)
		g.PrintBoard()
		hasWinner = g.isWin

		if !hasWinner {
			i++
		}
	}

	if hasWinner {
		fmt.Println("Player", i%2, "has won!")
	} else {
		fmt.Println("It's a tie!")
	}

	g.resetBoard()
}
