package board

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameMethods(t *testing.T) {
	var testCases = []struct {
		Name          string
		GameSetup     func() *Game
		TestMethod    func(g *Game) error
		ExpectedError error
		ExpectedState int
		Message       string
	}{
		{
			Name: "Test case 1: Open valid cell",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				return g
			},
			TestMethod: func(g *Game) error {
				return g.OpenCell(5, 5)
			},
			ExpectedError: nil,
			ExpectedState: Playing,
			Message:       "Opening a valid cell should not result in an error or game over, if not a mine",
		},
		{
			Name: "Test case 2: Open already opened cell",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				g.OpenCell(5, 5)
				return g
			},
			TestMethod: func(g *Game) error {
				return g.OpenCell(5, 5)
			},
			ExpectedError: nil,
			ExpectedState: Playing,
			Message:       "Opening already opened cell should not result in an error or change game state",
		},
		{
			Name: "Test case 3: Open cell outside the board",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				return g
			},
			TestMethod: func(g *Game) error {
				return g.OpenCell(15, 15)
			},
			ExpectedError: ErrInvalidCoordinates,
			ExpectedState: Playing,
			Message:       "Opening cell outside the board should result in ErrInvalidCoordinates error",
		},
		{
			Name: "Test case 4: Flag valid cell",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				return g
			},
			TestMethod: func(g *Game) error {
				return g.FlagCell(5, 5)
			},
			ExpectedError: nil,
			ExpectedState: Playing,
			Message:       "Flagging a valid cell should not result in an error",
		},
		{
			Name: "Test case 5: Flag cell outside the board",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				return g
			},
			TestMethod: func(g *Game) error {
				return g.FlagCell(15, 15)
			},
			ExpectedError: ErrInvalidCoordinates,
			ExpectedState: Playing,
			Message:       "Flagging cell outside the board should result in ErrInvalidCoordinates error",
		},
		{
			Name: "Test case 6: Unflag valid cell",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				g.FlagCell(5, 5)
				return g
			},
			TestMethod: func(g *Game) error {
				return g.UnflagCell(5, 5)
			},
			ExpectedError: nil,
			ExpectedState: Playing,
			Message:       "Unflagging a valid flagged cell should not result in an error",
		},
		{
			Name: "Test case 7: Unflag cell outside the board",
			GameSetup: func() *Game {
				g, _ := NewGame(10, 10, 10, "player")
				return g
			},
			TestMethod: func(g *Game) error {
				return g.UnflagCell(15, 15)
			},
			ExpectedError: ErrInvalidCoordinates,
			ExpectedState: Playing,
			Message:       "Unflagging cell outside the board should result in ErrInvalidCoordinates error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			game := tc.GameSetup()
			err := tc.TestMethod(game)
			require.Equal(t, tc.ExpectedError, err, tc.Message)
			require.Equal(t, tc.ExpectedState, game.GameState, "Game state mismatch")
		})
	}
}
