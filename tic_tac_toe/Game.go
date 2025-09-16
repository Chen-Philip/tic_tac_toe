package tic_tac_toe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	isGameOver bool
	turn       int
	board      Board
}

func (g *Game) getMove() (int, int) {
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
			} else if g.board.isValidMove(num2, num1) {
				x = num1
				y = num2
			}
		}
	}
	return y, x
}

func (g *Game) StartGame() {

	hasWinner := false
	i := 0
	for !hasWinner && i < 9 {
		x, y := g.getMove()
		g.board.makeMove(x, y, i%2 == 0)
		hasWinner = g.board.isWin
		if !hasWinner {
			i++
		}
	}

	if hasWinner {
		fmt.Println("Player", i%2, "has won!")
	} else {
		fmt.Println("It's a tie!")
	}
}
