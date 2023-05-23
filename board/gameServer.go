package board

import (
	"math/rand"
	"sync"
)

// GameServer responsible of store all gamse to the map
type GameServer struct {
	games map[string]*Game
	mu    sync.Mutex
}

func NewGameServer() *GameServer {
	return &GameServer{
		games: make(map[string]*Game),
	}
}

// NewGame creates a new game with the given parameters
func NewGame(width, height, mines int, player string) (*Game, error) {
	if width <= 0 || height <= 0 {
		return nil, ErrInvalidBoardSize
	}
	if mines <= 0 || mines >= width*height {
		return nil, ErrInvalidMineNumber
	}

	board := generateBoard(width, height, mines)
	stateBoard := make([][]int, height)
	for i := range stateBoard {
		stateBoard[i] = make([]int, width)
	}

	return &Game{
		Board:      board,
		StateBoard: stateBoard,
		Width:      width,
		Height:     height,
		Player:     player,
	}, nil
}

// generateBoard generates a new board with the given parameters
func generateBoard(width, height, mines int) [][]int {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}

	positions := make([]int, width*height)
	for i := range positions {
		positions[i] = i
	}
	rand.Shuffle(len(positions), func(i, j int) { positions[i], positions[j] = positions[j], positions[i] })

	for _, pos := range positions[:mines] {
		x, y := pos%width, pos/width
		board[y][x] = -1

		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				nx, ny := x+dx, y+dy
				if 0 <= nx && nx < width && 0 <= ny && ny < height && board[ny][nx] != -1 {
					board[ny][nx]++
				}
			}
		}
	}

	return board
}
