package tic_tac_toe

type Game struct {
	Board      [3][3]int
	winTracker [8]int
	IsWin      bool
	Turn       int
}

func (g *Game) ResetBoard() {
	// Resets the Board
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			g.Board[i][j] = 0
		}
	}
	// Resets the tracker
	for i := 0; i < len(g.winTracker); i++ {
		g.winTracker[i] = 0
	}
	// Resets the win flag
	g.IsWin = false
}

func (g *Game) MakeMove(x, y int) {
	player := 1
	if g.Turn%2 != 0 {
		player = -1
	}

	g.Turn += 1

	// Updates the tracker
	isWin := g.updateTracker(x, player)
	isWin = isWin || g.updateTracker(3+y, player)
	if x+y == 2 {
		isWin = isWin || g.updateTracker(6, player)
	}
	if x == y {
		isWin = isWin || g.updateTracker(7, player)
	}

	g.Board[x][y] = player
	g.IsWin = isWin
}

func (g *Game) updateTracker(i, player int) bool {
	g.winTracker[i] += player
	return g.winTracker[i] == 3 || g.winTracker[i] == -3
}

// Assumes the move is on the Board
func (g *Game) IsValidMove(x, y int) bool {
	return g.Board[x][y] == 0
}
