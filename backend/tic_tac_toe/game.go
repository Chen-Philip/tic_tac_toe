package tic_tac_toe

type Game struct {
	board      [3][3]int
	winTracker [8]int
	isWin      bool
}

func (g *Game) resetBoard() {
	// Resets the board
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {
			g.board[i][j] = 0
		}
	}
	// Resets the tracker
	for i := 0; i < len(g.winTracker); i++ {
		g.winTracker[i] = 0
	}
	// Resets the win flag
	g.isWin = false
}

func (g *Game) makeMove(x, y int, isPlayer1 bool) {
	player := 1
	if !isPlayer1 {
		player = -1
	}

	// Updates the tracker
	isWin := g.updateTracker(x, player)
	isWin = isWin || g.updateTracker(3+y, player)
	if x+y == 2 {
		isWin = isWin || g.updateTracker(6, player)
	}
	if x == y {
		isWin = isWin || g.updateTracker(7, player)
	}

	g.board[x][y] = player
	g.isWin = isWin
}

func (g *Game) updateTracker(i, player int) bool {
	g.winTracker[i] += player
	return g.winTracker[i] == 3 || g.winTracker[i] == -3
}

// Assumes the move is on the board
func (g *Game) isValidMove(x, y int) bool {
	return g.board[x][y] != 0
}
