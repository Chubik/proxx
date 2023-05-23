package board

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Default values for game parameters
const (
	defaultWidth  = 10
	defaultHeight = 10
	defaultMines  = 10

	// errors text constants
	ERROR_WRONG_COORDS_FORMAT = "wrong coords format"
	ERROR_GAME_NOT_FOUND      = "game not found"
)

// StatusHandler is an HTTP handler that displays the current status of the game.
// It takes the player's unique ID as a URL parameter and renders the game state in an HTML template.
// If the game is not found, it returns an HTTP 404 response.
func (gs *GameServer) StatusHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join("assets", "status.html"))
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, game); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// OpenCellHandler is an HTTP handler that opens a cell on the game board.
// It takes the player's unique ID and the x and y coordinates of the cell as URL parameters.
// If the game is not found or the coordinates are invalid, it returns an appropriate HTTP response.
// After the cell is opened, it redirects the client to the game status page.
func (gs *GameServer) OpenCellHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")
	xStr := chi.URLParam(r, "x")
	yStr := chi.URLParam(r, "y")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	if err := game.OpenCell(x, y); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}

// FlagCellHandler is an HTTP handler that flags a cell on the game board.
// It takes the player's unique ID and the x and y coordinates of the cell as URL parameters.
// If the game is not found or the coordinates are invalid, it returns an appropriate HTTP response.
// After the cell is flagged, it redirects the client to the game status page.
func (gs *GameServer) FlagCellHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")
	xStr := chi.URLParam(r, "x")
	yStr := chi.URLParam(r, "y")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	if err := game.FlagCell(x, y); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}

// UnflagCellHandler is an HTTP handler that unflags a cell on the game board.
// It takes the player's unique ID and the x and y coordinates of the cell as URL parameters.
// If the game is not found or the coordinates are invalid, it returns an appropriate HTTP response.
// After the cell is unflagged, it redirects the client to the game status page.
func (gs *GameServer) UnflagCellHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")
	xStr := chi.URLParam(r, "x")
	yStr := chi.URLParam(r, "y")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	if err := game.UnflagCell(x, y); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}

// StartGameHandler is an HTTP handler that starts a new game.
// It generates a unique player ID and creates a new game with default parameters.
// The new game is stored on the server and the client is redirected to the game status page.
func (gs *GameServer) StartGameHandler(w http.ResponseWriter, r *http.Request) {
	// generate unique player id
	player := uuid.NewString()

	// game creation
	game, err := NewGame(defaultWidth, defaultHeight, defaultMines, player)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add game to the server storage
	gs.mu.Lock()
	gs.games[player] = game
	gs.mu.Unlock()

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}
