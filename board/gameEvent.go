package board

import "errors"

// Cell statuses
const (
	Covered = iota
	Opened
	Flagged
)

// Game states
const (
	Playing = iota
	Won
	Lost
)

// Structure Game shows game status and information
type Game struct {
	Board      [][]int
	StateBoard [][]int
	Width      int
	Height     int
	GameState  int
	Player     string
}

// Errors that can be returned by the game
var (
	ErrInvalidBoardSize   = errors.New("size of the board should be positive")
	ErrInvalidMineNumber  = errors.New("count of mines should be positive and less than board size")
	ErrInvalidCoordinates = errors.New("coords should be positive and less than board size")
)

// OpenCell opens cell with coordinates x, y
func (g *Game) OpenCell(x, y int) error {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return ErrInvalidCoordinates
	}
	if g.GameState != Playing {
		return ErrInvalidCoordinates
	}
	if g.StateBoard[y][x] != Covered {
		return nil
	}

	g.StateBoard[y][x] = Opened
	switch g.Board[y][x] {
	case -1:
		g.GameState = Lost
	case 0:
		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				nx, ny := x+dx, y+dy
				if 0 <= nx && nx < g.Width && 0 <= ny && ny < g.Height {
					g.OpenCell(nx, ny)
				}
			}
		}
	}
	g.checkWin()

	return nil
}

// FlagCell sets flag on cell with coordinates x, y
func (g *Game) FlagCell(x, y int) error {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return ErrInvalidCoordinates
	}
	if g.GameState != Playing {
		return ErrInvalidCoordinates
	}

	g.StateBoard[y][x] = Flagged
	return nil
}

// UnflagCell unset flag on cell with coordinates x, y
func (g *Game) UnflagCell(x, y int) error {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return ErrInvalidCoordinates
	}
	if g.GameState != Playing {
		return ErrInvalidCoordinates
	}

	g.StateBoard[y][x] = Covered
	return nil
}

func (g *Game) checkWin() {
	for i := range g.Board {
		for j := range g.Board[i] {
			if g.Board[i][j] != -1 && g.StateBoard[i][j] != Opened && g.StateBoard[i][j] != Flagged {
				// If any safe cell is not opened or flagged, the game isn't won yet
				return
			}
		}
	}

	// If we made it through the whole board without returning,
	// then all safe cells are open or flagged, and the game is won
	g.GameState = Won
}
