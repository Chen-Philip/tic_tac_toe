package tic_tac_toe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BoardState struct {
	isGameOver bool
	turn       int
}

func getMove() (int, int) {
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
			} else {
				x = num1
				y = num2
			}
		}
	}
	return x, y
}

func (s *BoardState) StartGame() {
	i := 0
	for i < 3 {
		getMove()
		i++
	}
}
