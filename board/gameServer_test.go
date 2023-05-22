package board

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewGame(t *testing.T) {
	testCases := []struct {
		name    string
		width   int
		height  int
		mines   int
		err     error
		message string
	}{
		{
			name:    "Test case 1: Valid input",
			width:   10,
			height:  10,
			mines:   5,
			err:     nil,
			message: "Valid input parameters should not return an error",
		},
		{
			name:    "Test case 2: Invalid board size",
			width:   -1,
			height:  10,
			mines:   5,
			err:     ErrInvalidBoardSize,
			message: "Negative board dimensions should return ErrInvalidBoardSize",
		},
		{
			name:    "Test case 3: Invalid mine count",
			width:   10,
			height:  10,
			mines:   101,
			err:     ErrInvalidMineNumber,
			message: "Mine count greater than total cells should return ErrInvalidMineNumber",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewGame(tc.width, tc.height, tc.mines, "player")
			assert.Equal(t, tc.err, err, tc.message)
		})
	}
}
