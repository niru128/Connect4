package main

func findWinningMove(g *Game, symbol int) int {
	for c := 0; c < COLS; c++ {

		tmp := *g
		tmp.Turn = symbol

		if _, ok := tmp.DropDisc(c); ok && tmp.CheckWin(symbol) {
			return c
		}
	}
	return -1
}

func BotMove(g *Game, botSymbol int) int {

	opp := 1
	if botSymbol == 1 {
		opp = 2
	}

	// 1️⃣ Bot tries to WIN
	if c := findWinningMove(g, botSymbol); c != -1 {
		return c
	}

	// 2️⃣ Bot blocks opponent
	if c := findWinningMove(g, opp); c != -1 {
		return c
	}

	// 3️⃣ Center preference
	center := COLS / 2
	if g.Board[0][center] == 0 {
		return center
	}

	// 4️⃣ First valid
	for c := 0; c < COLS; c++ {
		if g.Board[0][c] == 0 {
			return c
		}
	}

	return -1
}
