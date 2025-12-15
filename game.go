package main

const (
	ROWS = 6
	COLS = 7
)

type Game struct {
	Board   [ROWS][COLS]int
	Turn    int
	Players map[int]*Player
}

// DropDisc attempts to drop a disc into the specified column
func (g *Game) DropDisc(col int) (int, bool) {

	if col < 0 || col >= COLS {
		return -1, false
	}

	for row := ROWS - 1; row >= 0; row-- {

		if g.Board[row][col] == 0 {
			g.Board[row][col] = g.Turn
			return row, true
		}
	}

	return -1, false
}

// CheckDirection checks if there are 4 in a row in a specific direction
func CheckDirection(
	board [ROWS][COLS]int,
	row, col,
	dr, dc,
	player int,

) bool {
	count := 0

	for i := 0; i < 4; i++ {
		r := row + dr*i
		c := col + dc*i

		if r < 0 || r >= ROWS || c < 0 || c >= COLS {
			return false
		}

		if board[r][c] != player {
			return false
		}

		count++
	}

	return count == 4
}

func (g *Game) CheckWin(player int) bool {
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {

			if g.Board[row][col] != player {
				continue
			}

			if CheckDirection(g.Board, row, col, 0, 1, player) || // horizontal
				CheckDirection(g.Board, row, col, 1, 0, player) || // vertical
				CheckDirection(g.Board, row, col, 1, 1, player) || // diagonal
				CheckDirection(g.Board, row, col, 1, -1, player) { // diagonal (towards left)
				return true
			}

		}
	}

	return false

}

func (g *Game) IsDraw() bool {
	for c := 0; c < COLS; c++ {
		if g.Board[0][c] == 0 {
			return false
		}
	}
	return true
}

func (g *Game) SwitchTurn() {
	if g.Turn == 1 {
		g.Turn = 2
	} else {
		g.Turn = 1
	}
}
